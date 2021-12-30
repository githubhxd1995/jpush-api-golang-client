package push

type Message struct {
	MsgContent  string                 `json:"msg_content"`            // 消息内容本身
	Title       string                 `json:"title,omitempty"`        // 消息标题
	ContentType string                 `json:"content_type,omitempty"` // 消息内容类型
	Extras      map[string]interface{} `json:"extras,omitempty"`       // JSON 格式的可选参数
}

// SetContent 设置消息内容
func (m *Message) SetContent(content string) {
	m.MsgContent = content
}

// SetTitle 设置消息标题
func (m *Message) SetTitle(title string) {
	m.Title = title
}

// SetContentType 设置消息内容类型
func (m *Message) SetContentType(contentType string) {
	m.ContentType = contentType
}

// SetExtras 设置消息扩展内容
func (m *Message) SetExtras(extras map[string]interface{}) {
	m.Extras = extras
}

// AddExtras 添加消息扩展内容
func (m *Message) AddExtras(key string, value interface{}) {
	if m.Extras == nil {
		m.Extras = make(map[string]interface{})
	}
	m.Extras[key] = value
}