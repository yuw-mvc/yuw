package configs

const (
	/**
	 * Yuw
	 */
	YuwVersion string = "version:1.0.0 -->"
	Yuw string = YuwVersion + `
 _       _               _             _   
(_)_   _(_)_         _  (_)           (_)  
  (_)_(_) (_)       (_) (_)     _     (_)  
    (_)   (_)       (_) (_)   _(_)_   (_) 
    (_)   (_)_  _  _(_)_(_)_(_)   (_)_(_)  
    (_)     (_)(_)(_) (_) (_)       (_)`


	/**
	 * Time Zone & Area Position
	 */
	LocationAsiaShanghai string = "Asia/Shanghai"
)

type (
	ResponsePoT struct {
		Status int
		Msg string
		Data interface{}
	}
)