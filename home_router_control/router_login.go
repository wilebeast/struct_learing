package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"
)

const legacyKey = "RDpbLfCPsJZ7fiv"

const legacyDict = "" +
	"yLwVl0zKqws7LgKPRQ84Mdt708T1qQ3Ha7xv3H7NyU84p21BriUWBU43odz3iP4r" +
	"BL3cD02KZciXTysVXiV8ngg6vL48rPJyAUw0HurW20xqxv9aYb4M9wK1Ae0wlro" +
	"510qXeU07kV57fQMc8L6aLgMLwygtc0F10a0Dg70TOoouyFhdysuRMO51yY5ZlO" +
	"ZZLEal1h0t9YQW0Ko7oBwmCAHoic4HYbUyVeU3sfQ1xtXcPcf1aT303wAQhv66q" +
	"zW"

type loginRequest struct {
	Method string       `json:"method"`
	Login  loginPayload `json:"login"`
}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	ErrorCode int    `json:"error_code"`
	Stok      string `json:"stok"`
}

type blacklistRequest struct {
	CustomNetwork customNetwork `json:"custom_network"`
	Method        string        `json:"method"`
}

type customNetwork struct {
	Name  string            `json:"name"`
	Table string            `json:"table"`
	Para  customNetworkPara `json:"para"`
}

type customNetworkPara struct {
	Hostname    string `json:"hostname"`
	MAC         string `json:"mac"`
	Enable      string `json:"enable"`
	Owner       string `json:"owner"`
	AccessRight string `json:"accessright"`
}

type blacklistExample struct {
	Name     string
	Hostname string
	MAC      string
	Enable   string
	Owner    string
	Access   string
}

var blacklistExamples = []blacklistExample{
	{
		Name:     "redmi-14c",
		Hostname: "Redmi-14C",
		MAC:      "A2-0F-64-9A-8C-88",
		Enable:   "on",
		Owner:    "web",
		Access:   "000",
	},
	{
		Name:     "lenovo-pc",
		Hostname: "lenovol%20pc",
		MAC:      "1C-3E-84-EB-E7-61",
		Enable:   "on",
		Owner:    "web",
		Access:   "000",
	},
	{
		Name:     "hasee-laptop",
		Hostname: "%E7%A5%9E%E8%88%9F%E7%AC%94%E8%AE%B0%E6%9C%AC",
		MAC:      "48-8A-D2-58-35-ED",
		Enable:   "on",
		Owner:    "web",
		Access:   "000",
	},
}

const (
	blacklistActionAdd    = "add"
	blacklistActionDelete = "delete"
	defaultInterval       = 15 * time.Minute
)

func encodePassword(password string) string {
	width := len(legacyKey)
	if len(password) > width {
		width = len(password)
	}

	var b strings.Builder
	b.Grow(width)
	for i := 0; i < width; i++ {
		left := 187
		right := 187
		if i < len(password) {
			left = int(password[i])
		}
		if i < len(legacyKey) {
			right = int(legacyKey[i])
		}
		b.WriteByte(legacyDict[(left^right)%len(legacyDict)])
	}
	return b.String()
}

func requestHeaders(base string) map[string]string {
	return map[string]string{
		"Host":             strings.TrimPrefix(strings.TrimPrefix(base, "http://"), "https://"),
		"X-Requested-With": "XMLHttpRequest",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/145.0.0.0 Safari/537.36",
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"DNT":              "1",
		"Content-Type":     "application/json; charset=UTF-8",
		"Origin":           base,
		"Referer":          base + "/",
		"Accept-Language":  "en,zh-CN;q=0.9,zh;q=0.8",
	}
}

func postJSON(client *http.Client, targetURL string, payload any, headers map[string]string) ([]byte, int, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewReader(body))
	if err != nil {
		return nil, 0, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}

func login(client *http.Client, baseURL, username, password string, passwordIsEncoded bool) (string, error) {
	loginPassword := password
	if !passwordIsEncoded {
		loginPassword = password
	}

	loginReq := loginRequest{
		Method: "do",
		Login: loginPayload{
			Username: username,
			Password: loginPassword,
		},
	}

	loginBody, statusCode, err := postJSON(client, baseURL+"/", loginReq, requestHeaders(baseURL))
	if err != nil {
		return "", fmt.Errorf("login request failed: %w", err)
	}

	fmt.Printf("[%s] login HTTP %d\n", time.Now().Format(time.RFC3339), statusCode)
	fmt.Println(string(loginBody))

	var loginResp loginResponse
	if err := json.Unmarshal(loginBody, &loginResp); err != nil {
		return "", fmt.Errorf("parse login response failed: %w", err)
	}

	if loginResp.ErrorCode != 0 {
		return "", fmt.Errorf("login failed: error_code=%d", loginResp.ErrorCode)
	}
	if loginResp.Stok == "" {
		return "", fmt.Errorf("login succeeded but no stok was returned")
	}

	return loginResp.Stok, nil
}

func blacklistName(mac string) string {
	replacer := strings.NewReplacer("-", "", ":", "", ".", "")
	return "black_" + strings.ToLower(replacer.Replace(mac))
}

func buildBlacklistRequest(example blacklistExample, action string) blacklistRequest {
	request := blacklistRequest{
		Method: action,
		CustomNetwork: customNetwork{
			Name: blacklistName(example.MAC),
		},
	}

	if action == blacklistActionAdd {
		request.CustomNetwork.Table = "mac_filter_black_entry"
		request.CustomNetwork.Para = customNetworkPara{
			Hostname:    example.Hostname,
			MAC:         strings.ToUpper(example.MAC),
			Enable:      example.Enable,
			Owner:       example.Owner,
			AccessRight: example.Access,
		}
	}

	return request
}

func applyBlacklistAction(client *http.Client, baseURL, stok, action string) error {
	blockURL := fmt.Sprintf("%s/stok=%s/ds", baseURL, stok)

	for _, example := range blacklistExamples {
		body, status, err := postJSON(client, blockURL, buildBlacklistRequest(example, action), requestHeaders(baseURL))
		if err != nil {
			return fmt.Errorf("%s %s failed: %w", action, example.Name, err)
		}

		fmt.Printf("[%s] %s %s HTTP %d\n", time.Now().Format(time.RFC3339), action, example.Name, status)
		fmt.Println(string(body))
	}

	return nil
}

func main() {
	baseURL := flag.String("base-url", "http://192.168.2.1", "router base URL")
	username := flag.String("username", "admin", "router username")
	password := flag.String("password", "", "router admin password")
	passwordIsEncoded := flag.Bool("password-is-encoded", false, "treat --password as already encoded")
	interval := flag.Duration("interval", defaultInterval, "time to keep devices in or out of blacklist per phase")
	timeout := flag.Duration("timeout", 10*time.Second, "HTTP timeout")
	flag.Parse()

	if *password == "" {
		fmt.Fprintln(os.Stderr, "missing required flag: --password")
		os.Exit(2)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	client := &http.Client{
		Jar:     jar,
		Timeout: *timeout,
	}

	base := strings.TrimRight(*baseURL, "/")

	if len(blacklistExamples) == 0 {
		fmt.Fprintln(os.Stderr, "no blacklist examples configured")
		os.Exit(1)
	}

	for {
		stok, err := login(client, base, *username, *password, *passwordIsEncoded)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("[%s] stok=%s\n", time.Now().Format(time.RFC3339), stok)

		if err := applyBlacklistAction(client, base, stok, blacklistActionDelete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("[%s] devices are now out of blacklist, sleeping for %s\n", time.Now().Format(time.RFC3339), *interval)
		time.Sleep(*interval)

		stok, err = login(client, base, *username, *password, *passwordIsEncoded)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("[%s] stok=%s\n", time.Now().Format(time.RFC3339), stok)

		if err := applyBlacklistAction(client, base, stok, blacklistActionAdd); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("[%s] devices are now in blacklist, sleeping for %s\n", time.Now().Format(time.RFC3339), *interval)
		time.Sleep(*interval)
	}
}
