<h1>submit:</h1>
<form method="POST" action='{{urlfor "SubmitController.Submit"}}'>
  {{ .xsrfdata }}
  {{template "alert.tpl" .}}

  <label for="inputTitle">Title</label>
  <div>
    <input placeholder="title" name="Title" value="{{index .Params " Title "}}" type="text" id="inputTitle" />
  </div>

  <label for="inputContent">Text</label>
  <div>
    <textarea placeholder="content" name="Text" value="{{index .Params " Content "}}" type="text" id="inputContent" ></textarea>
  </div>

  <label for="inputTopic">Topic</label>
  <div>
    <input placeholder="topic" name="Topic" value="{{index .Params " Topic "}}" type="text" id="inputTopic" />
  </div>

  <input type="submit" value="Submit">
</form>