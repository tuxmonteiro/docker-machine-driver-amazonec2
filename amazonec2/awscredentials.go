/*
 * Copyright (c) 2019 Globo.com
 * All rights reserved.
 *
 * This source is subject to the Apache License, Version 2.0.
 * Please see the LICENSE file for more information.
 *
 * Authors: See AUTHORS file
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package amazonec2

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsCredentials interface {
	Credentials() *credentials.Credentials
}

type ProviderFactory interface {
	NewStaticProvider(id, secret, token string) credentials.Provider
}

type defaultAWSCredentials struct {
	AccessKey        string
	SecretKey        string
	SessionToken     string
	providerFactory  ProviderFactory
	fallbackProvider awsCredentials
}

func NewAWSCredentials(id, secret, token string) *defaultAWSCredentials {
	creds := defaultAWSCredentials{
		AccessKey:        id,
		SecretKey:        secret,
		SessionToken:     token,
		fallbackProvider: &AwsDefaultCredentialsProvider{},
		providerFactory:  &defaultProviderFactory{},
	}
	return &creds
}

func (c *defaultAWSCredentials) Credentials() *credentials.Credentials {
	providers := []credentials.Provider{}
	if c.AccessKey != "" && c.SecretKey != "" {
		providers = append(providers, c.providerFactory.NewStaticProvider(c.AccessKey, c.SecretKey, c.SessionToken))
	}
	if c.fallbackProvider != nil {
		fallbackCreds, err := c.fallbackProvider.Credentials().Get()
		if err == nil {
			providers = append(providers, &credentials.StaticProvider{Value: fallbackCreds})
		}
	}
	return credentials.NewChainCredentials(providers)
}

type AwsDefaultCredentialsProvider struct{}

func (c *AwsDefaultCredentialsProvider) Credentials() *credentials.Credentials {
	return session.New().Config.Credentials
}

type defaultProviderFactory struct{}

func (c *defaultProviderFactory) NewStaticProvider(id, secret, token string) credentials.Provider {
	return &credentials.StaticProvider{Value: credentials.Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    token,
	}}
}
