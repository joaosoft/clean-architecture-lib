package validator

func (v *Validator) newDefaultMiddleHandlers() map[string]middleTagHandler {
	return map[string]middleTagHandler{
		constTagValue:      v.validate_value,
		constTagNot:        v.validate_not,
		constTagOptions:    v.validate_options,
		constTagNotOptions: v.validate_not_options,
		constTagSize:       v.validate_size,
		constTagMin:        v.validate_min,
		constTagMax:        v.validate_max,
		constTagNotEmpty:   v.validate_not_empty,
		constTagIsEmpty:    v.validate_is_empty,
		constTagNotNull:    v.validate_not_null,
		constTagIsNull:     v.validate_is_null,
		constTagRegex:      v.validate_regex,
		constTagCallback:   v.validate_callback,
		constTagAlpha:      v.validate_alpha,
		constTagNumeric:    v.validate_numeric,
		constTagBool:       v.validate_bool,
		constTagPassword:   v.validate_password,
		constTagPrefix:     v.validate_prefix,
		constTagSuffix:     v.validate_suffix,
		constTagContains:   v.validate_contains,
		constTagUUID:       v.validate_uuid,
		constTagIp:         v.validate_ip,
		constTagIpV4:       v.validate_ipv4,
		constTagIpV6:       v.validate_ipv6,
		constTagBase64:     v.validate_base64,
		constTagEmail:      v.validate_email,
		constTagURL:        v.validate_url,
		constTagHex:        v.validate_hex,
		constTagFile:       v.validate_file,

		constTagSet:         v.validate_set,
		constTagSetEmpty:    v.validate_set_empty,
		constTagSetDistinct: v.validate_set_distinct,
		constTagSetTrim:     v.validate_set_trim,
		constTagSetTitle:    v.validate_set_title,
		constTagSetLower:    v.validate_set_lower,
		constTagSetUpper:    v.validate_set_upper,
		constTagSetKey:      v.validate_set_key,
		constTagSetSanitize: v.validate_set_sanitize,
		constTagSetMd5:      v.validate_set_md5,
		constTagSetRandom:   v.validate_set_random,
	}
}