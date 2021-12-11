// Copyright © 2021 zc2638 <zc2638@qq.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wrapper

import (
	"context"

	"github.com/sirupsen/logrus"
)

type key int

const (
	keyLog key = iota
)

func WithLog(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, keyLog, logger)
}

func ContextLog(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(keyLog)
	if logger == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	return logger.(*logrus.Entry)
}
