package main

import (
	"fmt"
)

// ConvertToGoType converts the c datatype to a golang datatype.
func ConvertToGoType(cType string) (string, error) {
	switch cType {
	case "...":
		return "...any", nil
	case "AudioCallback":
		return cType, nil
	case "AudioStream":
		return cType, nil
	case "AutomationEvent":
		return cType, nil
	case "AutomationEventList":
		return cType, nil
	case "AutomationEventList *":
		return "*AutomationEventList", nil
	case "BoundingBox":
		return cType, nil
	case "Camera":
		return cType, nil
	case "Camera *":
		return "*Camera", nil
	case "Camera2D":
		return cType, nil
	case "Camera3D":
		return cType, nil
	case "Color":
		return "color.RGBA", nil
	case "Color *":
		return "*color.RGBA", nil
	case "FilePathList":
		return cType, nil
	case "Font":
		return cType, nil
	case "GlyphInfo":
		return cType, nil
	case "GlyphInfo *":
		return "*GlyphInfo", nil
	case "Image":
		return cType, nil
	case "Image *":
		return "*Image", nil
	case "LoadFileDataCallback":
		return cType, nil
	case "LoadFileTextCallback":
		return cType, nil
	case "Material":
		return cType, nil
	case "Material *":
		return "*Material", nil
	case "Matrix":
		return cType, nil
	case "Mesh":
		return cType, nil
	case "Mesh *":
		return "*Mesh", nil
	case "Model":
		return cType, nil
	case "Model *":
		return "*Model", nil
	case "ModelAnimation":
		return cType, nil
	case "ModelAnimation *":
		return "*ModelAnimation", nil
	case "Music":
		return cType, nil
	case "NPatchInfo":
		return cType, nil
	case "Ray":
		return cType, nil
	case "RayCollision":
		return cType, nil
	case "Rectangle":
		return cType, nil
	case "Rectangle **":
		return "[]*Rectangle", nil
	case "RenderTexture2D":
		return cType, nil
	case "SaveFileDataCallback":
		return cType, nil
	case "SaveFileTextCallback":
		return cType, nil
	case "Shader":
		return cType, nil
	case "Sound":
		return cType, nil
	case "Texture2D":
		return cType, nil
	case "Texture2D *":
		return "*Texture2D", nil
	case "TextureCubemap":
		return cType, nil
	case "TraceLogCallback":
		return cType, nil
	case "Vector2":
		return cType, nil
	case "Vector2 *":
		return "*Vector2", nil
	case "Vector3":
		return cType, nil
	case "Vector3 *":
		return "*Vector3", nil
	case "Vector4":
		return cType, nil
	case "VrDeviceInfo":
		return cType, nil
	case "VrStereoConfig":
		return cType, nil
	case "Wave":
		return cType, nil
	case "Wave *":
		return "*Wave", nil
	case "bool":
		return cType, nil
	case "char":
		return "int8", nil
	case "char *":
		return "string", nil
	case "const GlyphInfo *":
		return "*GlyphInfo", nil
	case "const Matrix *":
		return "*Matrix", nil
	case "const char *":
		return "string", nil
	case "const char **":
		return "[]string", nil
	case "const int *":
		return "[]int32", nil
	case "const unsigned char *":
		return "[]byte", nil
	case "const void *":
		return "unsafe.Pointer", nil
	case "double":
		return "float64", nil
	case "float":
		return "float32", nil
	case "float *":
		return "[]float32", nil
	case "int":
		return "int32", nil
	case "int *":
		return "[]int32", nil
	case "long":
		return "int64", nil
	case "unsigned char *":
		return "[]byte", nil
	case "unsigned int":
		return "uint32", nil
	case "void":
		return "", nil
	case "void *":
		return "unsafe.Pointer", nil
	default:
		return "", fmt.Errorf("unknown type %s", cType)
	}
}
