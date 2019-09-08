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

//resource "tumblr_post_quote" "first_quote" {
//    blog   = "${var.blog}"
//    quote  = "First image applied by terraform-provider-tumblr"
//    source = "https://terraform-provider-for.tumblr.com/"
//    tags   = "terraform,terraform provider,tumblr,quote"
//}
