// Code generaTed by fileb0x at "2019-07-04 12:13:14.104858 +0800 CST m=+0.189198518" from config file "b0x.yaml" DO NOT EDIT.
// modified(2019-07-04 12:08:17.116623297 +0800 CST)
// original path: swagger-ui/dist/swagger-ui-standalone-preset.js

package swaggerFiles

import (
	"os"
)

// FileSwaggerUIStandalonePresetJs is "/swagger-ui-standalone-preset.js"

func init() {

	f, err := FS.OpenFile(CTX, "/swagger-ui-standalone-preset.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileSwaggerUIStandalonePresetJs)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}