// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package indexpage

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Head() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width,initial-scale=1\"><title>syum.ai</title><link rel=\"me\" href=\"https://mstdn.jp/@_syumai\"><script src=\"https://unpkg.com/@tailwindcss/browser@4\"></script><script>\nasync function main() {\n  const avatarEl = document.getElementById('avatar');\n  const shareButton = document.getElementById('shareButton');\n  const reloadButton = document.getElementById('reloadButton');\n\n  let avatarDataUrl = '';\n  let sharedUrl;\n\n  const setAvatarImage = async () => {\n    try {\n      const result = await fetch('/image/random');\n      if (avatarDataUrl !== '') {\n        URL.revokeObjectURL(avatarDataUrl);\n      }\n      avatarDataUrl = URL.createObjectURL(await result.blob());\n      sharedUrl = result.url;\n    } catch (e) {\n      console.error(e);\n      return;\n    }\n    avatarEl.src = avatarDataUrl;\n  };\n\n  avatarEl.addEventListener('load', (e) => {\n    e.target.classList.remove('is-invisible');\n  });\n\n  reloadButton.addEventListener('click', async (e) => {\n    await setAvatarImage();\n  });\n\n  if (window.navigator.share) {\n    shareButton.addEventListener('click', (e) => {\n      window.navigator.share({\n        title: '#わたしのシュウマイ',\n        text: '#わたしのシュウマイ',\n        url: sharedUrl,\n      });\n    })\n  }\n\n  await setAvatarImage();\n}\n\nwindow.addEventListener('DOMContentLoaded', main);\n    </script></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
