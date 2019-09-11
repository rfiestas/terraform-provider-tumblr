package tumblr

import "github.com/hashicorp/terraform/helper/schema"

var fieldsAllPosts = []string{"state", "tags", "date", "format", "slug"}

func blogPostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: descriptions["blog"],
	}
}

func statePostSchema() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Description:  descriptions["state"],
		ValidateFunc: validateState,
	}
}

func tagsPostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: descriptions["tags"],
	}
}

func tweetPostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: descriptions["tweet"],
		Removed:     "Pending to implement, default is none",
	}
}

func datePostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: descriptions["date"],
	}
}

func formatPostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: descriptions["format"],
		Removed:     "Pending to implement, default is html",
	}
}

func slugPostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: descriptions["slug"],
		Removed:     "Pending to implement, default is dynamic",
	}
}

func nativeInlineImagesPostSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Description: descriptions["native_inline_images"],
		Removed:     "Pending to implement, default is false",
	}
}