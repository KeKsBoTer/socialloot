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
  <form method="POST" action='{{urlfor "ApiController.Submit"}}' enctype="multipart/form-data">
    <div>
      <input autofocus placeholder="Insert a title" name="title" value="" type="text" id="inputTitle"/>
    </div>
    <div>
      {{if eq .Type "text"}}
      <textarea placeholder="Write your text here..." name="content" type="text" id="inputContent"></textarea>
       {{else if eq .Type "link"}}
      <input type="text" name="content" placeholder="Insert URL e.g. http://google.de" /> 
      {{else if eq .Type "image"}}
      <input type="file" name="content" title="Upload image" accept="image/*" /> 
      {{end}}
    </div>

    {{if .Topic}}
    <input name="topic" value="{{.Topic.Name}}" type="hidden" id="inputTopic" /> {{else}}
    <label for="inputTopic">Topic</label>
    <div>
      <input style="width:auto" name="topic" value="" type="text" id="inputTopic" />
    </div>
    {{end}}

    <input type="hidden" name="type" value="{{.Type}}" />

    <p class="message"></p>
    <input type="submit" value="Submit">
  </form>
</div>