package main

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	samure "github.com/Samudevv/samurai-render-go"
	samureGL "github.com/Samudevv/samurai-render-go/backends/opengl"
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	vertexShaderSource = `
	#version 410 core
	layout (location = 0) in vec2 aPos;
	void main() {
		gl_Position = vec4(aPos.x, aPos.y, 0.0, 1.0);
	}
	`

	fragmentShaderSource = `
	#version 410 core
	out vec4 FragColor;
	void main() {
		FragColor = vec4(1.0, 0.5, 0.2, 1.0);
	}
	`
)

type App struct {
	vao    uint32
	vbo    uint32
	vio    uint32
	shader uint32
}

func (a *App) OnEvent(ctx samure.Context, event interface{}) {
	switch event.(type) {
	case samure.EventPointerButton:
		ctx.SetRunning(false)
	}
}

func (a *App) OnRender(ctx samure.Context, layerSurface samure.LayerSurface, outputGeo samure.Rect) {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(a.shader)
	gl.BindVertexArray(a.vao)
	gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_INT, nil)
}

func (a *App) OnUpdate(ctx samure.Context, deltaTime float64) {
}

func main() {
	a := &App{}

	cfg := samure.CreateContextConfig(a)
	cfg.PointerInteraction = true
	cfg.GL.MajorVersion = 4
	cfg.GL.MinorVersion = 1
	cfg.GL.Samples = 4

	bkd := &samureGL.Backend{}
	ctx, err := samure.CreateContextWithBackend(cfg, bkd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer ctx.Destroy()

	if err := bkd.InitBindings(ctx, gl.Init); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version:", version)

	// Set this so that you are not confused if nothing shows up
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)
	gl.Enable(gl.DEPTH_CLAMP)

	// Initialise Vertex and Index buffers
	vertices := []float32{
		-0.5, -0.5,
		0.5, -0.5,
		0.0, 0.5,
	}
	indices := []uint32{
		0, 1, 2,
	}

	gl.GenVertexArrays(1, &a.vao)
	gl.BindVertexArray(a.vao)

	gl.GenBuffers(1, &a.vbo)
	gl.GenBuffers(1, &a.vio)

	gl.BindBuffer(gl.ARRAY_BUFFER, a.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(vertices[0])), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, a.vio)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(unsafe.Sizeof(indices[0])), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, int32(2*unsafe.Sizeof(vertices[0])), gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)

	// Build Shader
	compileShader := func(sourceCode string, shaderType uint32) (uint32, error) {
		shader := gl.CreateShader(shaderType)

		csource, free := gl.Strs(sourceCode)
		gl.ShaderSource(shader, 1, csource, nil)
		free()
		gl.CompileShader(shader)

		var status int32
		gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
		if status == gl.FALSE {
			var logLength int32
			gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

			logMessage := strings.Repeat("\x00", int(logLength+1))
			gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logMessage))

			return 0, fmt.Errorf("shader compilation error: %s", logMessage)
		}

		return shader, nil
	}

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		fmt.Fprintln(os.Stderr, "VERTEX:", err)
		os.Exit(1)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FRAGMENT:", err)
		os.Exit(1)
	}

	a.shader = gl.CreateProgram()
	gl.AttachShader(a.shader, vertexShader)
	gl.AttachShader(a.shader, fragmentShader)
	gl.LinkProgram(a.shader)

	var status int32
	gl.GetProgramiv(a.shader, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(a.shader, gl.INFO_LOG_LENGTH, &logLength)

		logMessage := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(a.shader, logLength, nil, gl.Str(logMessage))

		fmt.Fprintln(os.Stderr, "SHADER link error:", logMessage)
		os.Exit(1)
	}

	// Some GPUs need a vertex array bound while validating
	var tmpVao uint32
	gl.GenVertexArrays(1, &tmpVao)
	gl.BindVertexArray(tmpVao)
	gl.ValidateProgram(a.shader)
	gl.BindVertexArray(tmpVao)
	gl.DeleteVertexArrays(1, &tmpVao)

	gl.GetProgramiv(a.shader, gl.VALIDATE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(a.shader, gl.INFO_LOG_LENGTH, &logLength)

		logMessage := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(a.shader, logLength, nil, gl.Str(logMessage))

		fmt.Fprintln(os.Stderr, "SHADER validate error:", logMessage)
		os.Exit(1)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	ctx.Run()
}
