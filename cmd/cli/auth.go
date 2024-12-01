package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"ai-powered-study-planner-backend/pkg/auth"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

var GetJWTTokenCommand = cli.Command{
	Name:     "jwt",
	Usage:    "load user jwt by using user's UID",
	Args:     true,
	Category: "auth",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "uid",
			Aliases:     []string{"u"},
			Usage:       "firebase `UID` of the user",
			Value:       "YfA8G9E54aWvNajPmgdeyFmHI7N2",
			DefaultText: "UID",
		},
	},
	Action: func(ctx *cli.Context) error {
		uid := ctx.String("uid")
		if uid == "" {
			return fmt.Errorf("user's UID is required")
		}

		token, err := GetIdTokenByUID(uid)
		if err != nil {
			return cli.Exit(err, 1)
		}

		fmt.Println("Token:", token)

		return nil
	},
}

const baseUri = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s"

func GetIdTokenByUID(uid string) (string, error) {
	authClient := auth.GetFirebaseManager().Client

	// Create a custom token from the Firebase UID
	customToken, err := authClient.CustomToken(context.Background(), uid)
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("FIREBASE_WEB_API_KEY")
	url := fmt.Sprintf(baseUri, apiKey)
	requestBody := map[string]any{
		"token":             customToken,
		"returnSecureToken": true,
	}

	json, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return gjson.Get(string(body), "idToken").String(), nil
}
