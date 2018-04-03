<h1>submit:</h1>
<ul class="tab-list">
    <li>
      <a href="{{ChangeParam .URL "type" "link"}}" class="uppercase">link</a>
    </li>
    <li>
      <a href="{{ChangeParam .URL "type" "text"}}" class="uppercase">text</a>
    </li>
    <li>
      <a href="{{ChangeParam .URL "type" "image"}}" class="uppercase">image</a>
    </li>
</ul>
<form method="POST" action='{{urlfor "ApiController.Submit"}}' enctype="multipart/form-data">
  <label for="inputTitle">Title</label>
  <div>
    <input placeholder="title" name="title" value="" type="text" id="inputTitle" />
  </div>

  <label for="inputContent">{{.Type}}</label>
  <div>
  {{if eq .Type "text"}}
    <textarea placeholder="Write your text here..." name="content" type="text" id="inputContent"></textarea>
  {{else if eq .Type "link"}}
   <input type="text" name="content" placeholder="link e.g. google.de"/>
  {{else if eq .Type "image"}}
    <input type="file" name="content" title="Upload image" accept="image/*"/>
  {{end}}
  </div>
  
  {{if .Topic}}
    <input name="topic" value="{{.Topic.Name}}" type="hidden" id="inputTopic" />
  {{else}}
    <label for="inputTopic">Topic</label>
    <div>
      <input placeholder="topic" name="topic" value="" type="text" id="inputTopic" />
    </div>
  {{end}}
  
  <input type="hidden" name="type" value="{{.Type}}"/>

  <p class="message"></p>
  <input type="submit" value="Submit">
</form>