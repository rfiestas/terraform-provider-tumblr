resource "tumblr_post_text" "first_text" {
    blog  = "${var.blog}"
    title = "First text applied by terraform-provider-tumblr"
    body  = <<EOF
<p><pre class="language-terraform"><code>
resource "tumblr_post_text" "first_text" {
    blog  = "${var.blog}"
    title = "First text applied by terraform-provider-tumblr"
    body  = "First text applied by terraform-provider-tumblr"
    state = "published"
    tags  = "terraform,terraform provider,tumblr,text"
    date  = "2006-01-02 15:04:05"
}
</code></pre></p>
EOF
    state = "published"
    tags  = "terraform,terraform provider,tumblr,text"
    date  = "2006-01-02 15:04:05"
}

resource "tumblr_post_photo" "first_photo" {
    blog        = "${var.blog}"
    caption     = "First image applied by terraform-provider-tumblr"
    link        = "https://terraform-provider-for.tumblr.com/"
    data64_file = "/Users/rfiestas/Downloads/terraform-logo.png" 
    //state     = "published"
    tags        = "terraform,terraform provider,tumblr,photo"
    //date      = "2006-01-02 15:04:05"
}

resource "tumblr_post_quote" "first_quote" {
    blog   = "${var.blog}"
    quote  = "Anyone can develop and distribute their own Terraform providers."
    source = "https://www.terraform.io/docs/configuration/providers.html#third-party-plugins"
    tags   = "terraform,terraform provider,tumblr,quote"
}

resource "tumblr_post_link" "first_link" {
    blog        = "${var.blog}"
    title       = "terraform provider for tumblr"
    url         = "https://github.com/rfiestas/terraform-provider-tumblr"
    thumbnail   = "https://repository-images.githubusercontent.com/206975372/3a24af00-d275-11e9-8c20-11bd31b5aaaf"
    author      = "rfiestas"
    tags        = "terraform,terraform provider,tumblr,link"
}