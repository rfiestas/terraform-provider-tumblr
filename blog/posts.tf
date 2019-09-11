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
    data64      = "${filebase64("/Users/rfiestas/Downloads/terraform-logo.png")}"
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
    description = "The Tumblr provider is used to interact with the many resources supported by tumblr.com. The provider needs to be configured with the proper credentials before it can be used."
    thumbnail   = "https://repository-images.githubusercontent.com/206975372/3a24af00-d275-11e9-8c20-11bd31b5aaaf"
    tags        = "terraform,terraform provider,tumblr,link"
}


resource "tumblr_post_audio" "first_audio" {
    blog         = "${var.blog}"
    caption      = "First audio applied by terraform-provider-tumblr"    
    external_url = "https://soundcloud.com/club-bizarre/free-download-kraftwerk-we-are"
    tags         = "terraform,terraform provider,tumblr,audio"
}

resource "tumblr_post_video" "first_video" {
    blog         = "${var.blog}"
    caption      = "First video applied by terraform-provider-tumblr"    
    embed        = "https://www.youtube.com/watch?v=TMayVLSQ6yM"
    tags         = "terraform,terraform provider,tumblr,video"
}

resource "tumblr_post_chat" "first_chat" {
    blog         = "${var.blog}"
    title        = "First chat applied by terraform-provider-tumblr"    
    conversation = "${file("resources/chat.txt")}"
    tags         = "terraform,terraform provider,tumblr,chat"
}

resource "tumblr_post_photo" "apollo12" {
    blog        = "${var.blog}"
    caption     = "<p>50 Years Ago Apollo 11 Launches Into History</p><small>At 9:32 a.m. EDT, July 16, 1969, Apollo 11 launched from Florida on a mission to the Moon.</small>"
    link        = "https://www.nasa.gov/multimedia/imagegallery/iotd.html"
    source      = "https://www.nasa.gov/sites/default/files/styles/full_width_feature/public/thumbnails/image/liftoff_0.jpg"
    tags        = "terraform,terraform provider,tumblr,photo,nasa,apollo"
    
}