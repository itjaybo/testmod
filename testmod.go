package testmod

import (
    "errors"
    "fmt" 
) 

// Hi 返回一个 `lang` 语言的问候语
func Hi(name, lang string) (string, error) {
    switch lang {
    case "cn":
        return fmt.Sprintf("你好，%s！", name), nil
    case "en":
        return fmt.Sprintf("Hi, %s!", name), nil
    case "pt":
        return fmt.Sprintf("Oi, %s!", name), nil
    case "es":
        return fmt.Sprintf("¡Hola, %s!", name), nil
    case "fr":
        return fmt.Sprintf("Bonjour, %s!", name), nil
    default:
        return "", errors.New("未知的语言")
    }
}