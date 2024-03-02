package tmplUtil

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xml520/go-utils/strUtil"
	"github.com/xml520/go-utils/structUtil"
	"os"
	"strings"
)

const (
	defaultGroupName = "router"
)

var (
	renderStart = []byte("//** router **//")
)

type Router struct {
	st *structUtil.Struct
}
type routerBlock struct {
	level int
	end   bool
	Names structUtil.StructNames
}
type RouterMethod struct {
	Method  string
	Url     string
	Handler string
}

func NewRouter(v *structUtil.Struct) *Router {

	return &Router{st: v}
}

func (r *Router) RouterBlock() (blocks []routerBlock) {
	names := r.st.Names()
	for i, _ := range names {
		blocks = append(blocks, routerBlock{
			level: i,
			Names: names[:i+1],
			end:   len(names) == i+1,
		})
	}
	return blocks
}
func (r *Router) Render(raw []byte, method ...RouterMethod) ([]byte, error) {
	if len(raw) == 0 {
		raw = append(raw, renderStart...)
	}
	var err error
	blocks := r.RouterBlock()
	for _, block := range blocks {
		if !block.IsExists(raw) {
			if block.end {
				raw, err = block.Write(raw, method...)
			} else {
				raw, err = block.Write(raw)
			}

			if err != nil {
				return nil, err
			}
		}
	}
	return raw, err
}
func (r *Router) Write(filename string, method ...RouterMethod) error {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	raw, err = r.Render(raw, method...)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, raw, 0644)
}
func (r *routerBlock) Write(raw []byte, method ...RouterMethod) (b []byte, err error) {
	superMarkStr := r.SuperBlock().Mark()
	startIndex := bytes.Index(raw, []byte(superMarkStr))
	if startIndex == -1 {
		return nil, errors.New("找不到上级块标记 " + superMarkStr)
	}
	raw = append(raw[:startIndex], append([]byte(r.Render(method...)), raw[startIndex:]...)...)
	return raw, err
}
func (r *routerBlock) IsExists(raw []byte) bool {
	return bytes.Index(raw, []byte(r.Mark())) != -1
}
func (r *routerBlock) Render(methods ...RouterMethod) string {
	groupName := r.GroupName()
	common := fmt.Sprintf("// %sGroup", r.Names.String())
	superGroupName := r.SuperBlock().Names.LowerCamel()
	groupUri := strUtil.ToLower(r.Names[len(r.Names)-1])
	methodStr := r.RenderMethod(methods)
	mark := r.Mark()
	return fmt.Sprintf("\n%s\n%s := %s.Group(\"/%s\")\n%s\n%s\n%s\n%s\n%s",
		r.Sep(1, common),
		r.Sep(1, groupName),
		superGroupName,
		groupUri,
		r.Sep(1, "{"),
		methodStr,
		r.Sep(2, mark),
		r.Sep(1, "}"),
		r.Sep(1),
	)
}
func (r *routerBlock) RenderMethod(methods []RouterMethod) string {
	if len(methods) == 0 {
		return ""
	}
	var methodArr []string
	for _, method := range methods {
		groupName := r.GroupName()
		methodStr := strUtil.ToUpper(strUtil.ToLower(method.Method), 1)
		handler := strings.ReplaceAll(method.Handler, "*", r.Names.String())
		methodArr = append(methodArr, r.Sep(2, fmt.Sprintf("%s.%s(\"%s\",%s)", groupName, methodStr, "/"+method.Url, handler)))
	}
	return strings.Join(methodArr, "\n")
}
func (r *routerBlock) GroupName() string {
	return r.Names.LowerCamel()
}
func (r *routerBlock) Mark() string {
	return fmt.Sprintf("//** %s **//", r.Names.String())
}
func (r *routerBlock) Sep(n int, str ...any) string {
	return fmt.Sprintf("%s%s", strings.Repeat("\t", (r.level+1)*n), fmt.Sprint(str...))
}
func (r *routerBlock) SuperBlock() *routerBlock {
	if len(r.Names) == 1 {
		return &routerBlock{Names: structUtil.StructNames{defaultGroupName}, level: r.level - 1}
	} else {
		return &routerBlock{Names: r.Names[:len(r.Names)-1], level: r.level - 1}
	}
}
func (r *routerBlock) IsEnd() bool {
	return r.end
}
