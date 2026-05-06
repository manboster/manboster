package gguf

var models = []Model{
	{
		DisplayName: "Qwen3 Guard",
		Name:        "qwen3-guard",
		Description: "Qwen3Guard is a series of safety moderation models built upon Qwen3 and trained on a dataset of 1.19 million prompts and responses labeled for safety. ",
		Groups: []Group{
			{
				Parameters: "0.6B",
				Quants: []Quant{
					{
						DisplayName: "Q8_0 (Smartest but bigger and occupied more memory)",
						Mod:         "Q8_0",
						URL:         "https://huggingface.co/QuantFactory/Qwen3Guard-Gen-0.6B-GGUF/resolve/main/Qwen3Guard-Gen-0.6B.Q8_0.gguf",
						Sha256:      "a5108bc29824fbe805d4674cf8af438bb7221186c2fa940175a8dfdda4a8d2cc",
						Size:        "805M",
					},
					{
						DisplayName: "Q6_K",
						Mod:         "Q6_K",
						URL:         "https://huggingface.co/QuantFactory/Qwen3Guard-Gen-0.6B-GGUF/resolve/main/Qwen3Guard-Gen-0.6B.Q6_K.gguf",
						Sha256:      "33a70125c0fff6805e1a1b8b99f59981c6cd3f724a06bf3a99c16dfa8326e585",
						Size:        "623M",
					},
					{
						DisplayName: "Q5_K_M",
						Mod:         "Q5_K_M",
						URL:         "https://huggingface.co/QuantFactory/Qwen3Guard-Gen-0.6B-GGUF/resolve/main/Qwen3Guard-Gen-0.6B.Q5_K_M.gguf",
						Sha256:      "023c0136c0232e6b1028e2c88ef2b1f28a74938a720146ca0e13fb0b24f105d6",
						Size:        "551M",
					},
					{
						DisplayName: "Q4_K_M (balanced selection)",
						Mod:         "Q4_K_M",
						URL:         "https://huggingface.co/QuantFactory/Qwen3Guard-Gen-0.6B-GGUF/resolve/main/Qwen3Guard-Gen-0.6B.Q4_K_M.gguf",
						Sha256:      "aa114c4bfa17b0dca5e08b4d27ca21d7ba2f21c0423f5661e0f0caf799b3284c",
						Size:        "484M",
					},
					{
						DisplayName: "Q3_K_M",
						Mod:         "Q3_K_M",
						URL:         "https://huggingface.co/QuantFactory/Qwen3Guard-Gen-0.6B-GGUF/resolve/main/Qwen3Guard-Gen-0.6B.Q3_K_M.gguf",
						Sha256:      "52c71bc5ff115f553dc793a3ce8cc2916bc259135985e81a917575fb0c1d05da",
						Size:        "414M",
					},
					{
						DisplayName: "Q2_K (Smaller but not smart)",
						Mod:         "Q2_K",
						URL:         "https://huggingface.co/QuantFactory/Qwen3Guard-Gen-0.6B-GGUF/resolve/main/Qwen3Guard-Gen-0.6B.Q2_K.gguf",
						Sha256:      "aa114c4bfa17b0dca5e08b4d27ca21d7ba2f21c0423f5661e0f0caf799b3284c",
						Size:        "347M",
					},
				},
			},
			{
				Parameters: "4B",
				Quants:     []Quant{},
			},
			{
				Parameters: "8B",
				Quants:     []Quant{},
			},
		},
	},
}
