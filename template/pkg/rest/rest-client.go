package rest

import (
	"time"

	"{{bootstrap_template}}/pkg/log/logger"

	"github.com/go-resty/resty/v2"
)

type RestClient struct {
	*resty.Client
	logger.Logger
}

func newRestyV2(config *config, logger logger.Logger) *RestClient {
	c := newTemplate(config.timeout, logger)
	return &RestClient{
		c,
		logger,
	}
}

func (c *RestClient) Post(path string, body interface{}, headers map[string]string) (int, []byte, error) {
	response, err := c.R().
		SetHeaders(headers).
		SetBody(body).
		Post(path)
	return response.StatusCode(), response.Body(), err
}

func (c *RestClient) Get(path string, headers map[string]string) (int, []byte, error) {
	response, err := c.R().
		SetHeaders(headers).
		Get(path)
	return response.StatusCode(), response.Body(), err
}

func (c *RestClient) GetWithParams(path string, headers map[string]string, params map[string]string) (int, []byte, error) {
	response, err := c.R().
		SetHeaders(headers).
		SetQueryParams(params).
		Get(path)
	return response.StatusCode(), response.Body(), err
}

func newTemplate(timeout int, logger logger.Logger) *resty.Client {
	c := resty.New()
	c.SetRetryCount(1).SetLogger(logger).SetTimeout(time.Duration(timeout) * time.Second)
	return c
}
