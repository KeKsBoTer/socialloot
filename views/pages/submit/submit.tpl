<div class="auth-form full-width">
  <h1>submit</h1>
  <ul class="tab-list">
    <li class="{{if eq (GetParam .URL "type") "link"}} active{{end}}">
      <a href="{{ChangeParam .URL "type" "link"}}" class="uppercase">link</a>
    </li>
    <li class="{{if or (eq (GetParam .URL "type") "text") (eq (GetParam .URL "type") "")}} active{{end}}">
      <a href="{{ChangeParam .URL "type" "text"}}" class="uppercase">text</a>
    </li>
    <li class="{{if eq (GetParam .URL "type") "image"}} active{{end}}">
      <a href="{{ChangeParam .URL "type" "image"}}" class="uppercase">image</a>
    </li>
  </ul>
  <form method="POST" action='{{urlfor "APIController.Submit"}}' enctype="multipart/form-data">
    <div>
      <label for="inputTitle" class="message"></label>
      <input autofocus placeholder="Insert a title" name="title" value="" type="text" id="inputTitle"/>
    </div>
    <div>
      {{if eq .Type "text"}}
      <label for="inputContent" class="message"></label>
      <textarea placeholder="Write your text here..." name="content" type="text" id="inputContent"></textarea>
       {{else if eq .Type "link"}}
       <label for="insertLink" class="message"></label>
      <input type="text" name="content" placeholder="Insert URL e.g. http://google.de" / id="insertLink"> 
      {{else if eq .Type "image"}}
      <label for="insertImage" class="message"></label>
      <input type="file" name="content" id="insertImage" title="Upload image" accept="image/*" 
              onchange="previewImage(this,$('#image-preview'));"/> 
      <img id="image-preview" src="#"/>
      {{end}}
    </div>

    {{if .Topic}}
      <input name="topic" value="{{.Topic.Name}}" type="hidden" id="inputTopic" /> 
    {{else}}
    <label for="inputTopic">Topic</label>
    <div>
      <label for="inputTopic" class="message"></label>
      <input style="width:auto" name="topic" value="" type="text" id="inputTopic" />
    </div>
    {{end}}

    <input type="hidden" name="type" value="{{.Type}}" />

    <p class="global-message"></p>
    <input type="submit" value="Submit{{if .Topic}} to {{.Topic.Name}}{{end}}">
  </form>
</div>