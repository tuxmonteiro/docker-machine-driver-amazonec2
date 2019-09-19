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
	"errors"
)

type region struct {
	AmiId string
}

// Ubuntu 18.04 LTS 20190911 hvm:ebs-ssd (amd64)
// See https://cloud-images.ubuntu.com/locator/ec2/
var regionDetails map[string]*region = map[string]*region{
	"ap-east-1":       {"ami-d8cdb6a9"},
	"ap-northeast-1":  {"ami-0193f3ddafc2d8921"},
	"ap-south-1":      {"ami-03d48fdf7c142da15"},
	"ap-southeast-1":  {"ami-0f8c65cdcc35d0a72"},
	"ca-central-1":    {"ami-0e76f0489f3cca311"},
	"eu-central-1":    {"ami-0fd6d676e2e90dc7b"},
	"eu-north-1":      {"ami-dbc04aa5"},
	"eu-west-1":       {"ami-08f053fa3d25478f4"},
	"me-south-1":      {"ami-07ca5b987e63907ba"},
	"sa-east-1":       {"ami-0017950c58e65a5a4"},
	"us-east-1":       {"ami-024582e76075564db"},
	"us-west-1":       {"ami-0c6eca62e6c7dacff"},
	"cn-north-1":      {"ami-01993b4213b4bffb5"}, // Note: this is 20190627.1
	"cn-northwest-1":  {"ami-01d4e30d4d4952d0f"}, // Note: this is 20190627.1
	"us-gov-west-1":   {"ami-60501e01"},          // Note: this is 20190814
	"us-gov-east-1":   {"ami-cce606bd"},          // Note: this is 20190814
	"ap-northeast-2":  {"ami-0b0d9551c4775f4e4"},
	"ap-southeast-2":  {"ami-056280928f86f80a8"},
	"eu-west-2":       {"ami-012d95ddd53c98e0e"},
	"us-east-2":       {"ami-090a9429be63ca087"},
	"us-west-2":       {"ami-09c08e9cd8d9c1b93"},
	"ap-northeast-3":  {"ami-08801b4823195185b"},
	"eu-west-3":       {"ami-0001c7938b8e89312"},
	"custom-endpoint": {""},
}

func awsRegionsList() []string {
	var list []string

	for k := range regionDetails {
		list = append(list, k)
	}

	return list
}

func validateAwsRegion(region string) (string, error) {
	for _, v := range awsRegionsList() {
		if v == region {
			return region, nil
		}
	}

	return "", errors.New("Invalid region specified")
}
