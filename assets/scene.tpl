{{- $l := len . }}
{{- range $idx, $bat := . }}
    {{- color }}{{ indent $bat.Location $bat.Sprite }}{{ colorReset }}
    {{- if ne (add $idx 1) $l }}
        {{ batSpace }}
    {{- end }}
{{- end }}