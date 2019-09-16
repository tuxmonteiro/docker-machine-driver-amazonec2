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

package awsec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"os"
)

type awslogger struct {
	logger *log.Logger
}

func AwsLogger() aws.Logger {
	return &awslogger{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l awslogger) Log(args ...interface{}) {
	l.logger.Println(args...)
}
