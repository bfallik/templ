// Code generated by templ@(devel) DO NOT EDIT.

package testcssusage

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

func green() templ.CSSClass {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`color:#00ff00;`)
	templCSSID := templ.CSSID(`green`, templCSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}

func className() templ.CSSClass {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`background-color:#ffffff;`)
	templCSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, red)))
	templCSSID := templ.CSSID(`className`, templCSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}

func Button(text string) templ.Component {
	return templ.ComponentFunc(func(templCtx context.Context, templW io.Writer) (templErr error) {
		templBuffer, templIsBuffer := templW.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		templCtx = templ.InitializeContext(templCtx)
		templVar1 := templ.GetChildren(templCtx)
		if templVar1 == nil {
			templVar1 = templ.NopComponent
		}
		templCtx = templ.ClearChildren(templCtx)
		var templVar2 = []any{className(), templ.Class("&&&unsafe"), "safe", templ.SafeClass("safe2")}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar2...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<button class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar2).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\" type=\"button\">")
		if templErr != nil {
			return templErr
		}
		var templVar3 string = text
		_, templErr = templBuffer.WriteString(templ.EscapeString(templVar3))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("</button>")
		if templErr != nil {
			return templErr
		}
		if !templIsBuffer {
			_, templErr = templBuffer.WriteTo(templW)
		}
		return templErr
	})
}

func LegacySupport() templ.Component {
	return templ.ComponentFunc(func(templCtx context.Context, templW io.Writer) (templErr error) {
		templBuffer, templIsBuffer := templW.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		templCtx = templ.InitializeContext(templCtx)
		templVar4 := templ.GetChildren(templCtx)
		if templVar4 == nil {
			templVar4 = templ.NopComponent
		}
		templCtx = templ.ClearChildren(templCtx)
		var templVar5 = []any{templ.Classes(templ.Class("test"), "a")}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar5...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<div class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar5).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\"></div>")
		if templErr != nil {
			return templErr
		}
		if !templIsBuffer {
			_, templErr = templBuffer.WriteTo(templW)
		}
		return templErr
	})
}

func MapCSSExample() templ.Component {
	return templ.ComponentFunc(func(templCtx context.Context, templW io.Writer) (templErr error) {
		templBuffer, templIsBuffer := templW.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		templCtx = templ.InitializeContext(templCtx)
		templVar6 := templ.GetChildren(templCtx)
		if templVar6 == nil {
			templVar6 = templ.NopComponent
		}
		templCtx = templ.ClearChildren(templCtx)
		var templVar7 = []any{map[string]bool{"a": true, "b": false, "c": true}}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar7...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<div class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar7).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\"></div>")
		if templErr != nil {
			return templErr
		}
		if !templIsBuffer {
			_, templErr = templBuffer.WriteTo(templW)
		}
		return templErr
	})
}

func KVExample() templ.Component {
	return templ.ComponentFunc(func(templCtx context.Context, templW io.Writer) (templErr error) {
		templBuffer, templIsBuffer := templW.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		templCtx = templ.InitializeContext(templCtx)
		templVar8 := templ.GetChildren(templCtx)
		if templVar8 == nil {
			templVar8 = templ.NopComponent
		}
		templCtx = templ.ClearChildren(templCtx)
		var templVar9 = []any{"a", templ.KV("b", false)}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar9...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<div class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar9).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\"></div>")
		if templErr != nil {
			return templErr
		}
		var templVar10 = []any{"a", "b", "c", templ.KV("c", false)}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar10...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<input type=\"email\" id=\"email\" name=\"email\" class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar10).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\" placeholder=\"your@email.com\" autocomplete=\"off\">")
		if templErr != nil {
			return templErr
		}
		if !templIsBuffer {
			_, templErr = templBuffer.WriteTo(templW)
		}
		return templErr
	})
}

func PsuedoAttributes() templ.Component {
	return templ.ComponentFunc(func(templCtx context.Context, templW io.Writer) (templErr error) {
		templBuffer, templIsBuffer := templW.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		templCtx = templ.InitializeContext(templCtx)
		templVar11 := templ.GetChildren(templCtx)
		if templVar11 == nil {
			templVar11 = templ.NopComponent
		}
		templCtx = templ.ClearChildren(templCtx)
		var templVar12 = []any{"bg-violet-500", templ.KV(templ.SafeClass("hover:bg-violet-600"), true)}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar12...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<button class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar12).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\">")
		if templErr != nil {
			return templErr
		}
		templVar13 := `Save changes`
		_, templErr = templBuffer.WriteString(templVar13)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("</button>")
		if templErr != nil {
			return templErr
		}
		if !templIsBuffer {
			_, templErr = templBuffer.WriteTo(templW)
		}
		return templErr
	})
}

func ThreeButtons() templ.Component {
	return templ.ComponentFunc(func(templCtx context.Context, templW io.Writer) (templErr error) {
		templBuffer, templIsBuffer := templW.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		templCtx = templ.InitializeContext(templCtx)
		templVar14 := templ.GetChildren(templCtx)
		if templVar14 == nil {
			templVar14 = templ.NopComponent
		}
		templCtx = templ.ClearChildren(templCtx)
		templErr = Button("A").Render(templCtx, templBuffer)
		if templErr != nil {
			return templErr
		}
		templErr = Button("B").Render(templCtx, templBuffer)
		if templErr != nil {
			return templErr
		}
		var templVar15 = []any{templ.Classes(green)}
		templErr = templ.RenderCSSItems(templCtx, templBuffer, templVar15...)
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("<button class=\"")
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(templVar15).String()))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("\" type=\"button\">")
		if templErr != nil {
			return templErr
		}
		var templVar16 string = "Green"
		_, templErr = templBuffer.WriteString(templ.EscapeString(templVar16))
		if templErr != nil {
			return templErr
		}
		_, templErr = templBuffer.WriteString("</button>")
		if templErr != nil {
			return templErr
		}
		templErr = MapCSSExample().Render(templCtx, templBuffer)
		if templErr != nil {
			return templErr
		}
		templErr = KVExample().Render(templCtx, templBuffer)
		if templErr != nil {
			return templErr
		}
		templErr = PsuedoAttributes().Render(templCtx, templBuffer)
		if templErr != nil {
			return templErr
		}
		if !templIsBuffer {
			_, templErr = templBuffer.WriteTo(templW)
		}
		return templErr
	})
}
