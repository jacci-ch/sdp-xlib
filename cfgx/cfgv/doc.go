// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

// Module logx uses cfgx module to load logger configurations
// and the cfgx module uses logx to output info messages.
//
// So we need to extract a middleware module cfgv to use to avoid
// 'import cycle' error while compiling the code.

package cfgv
