<h1>Create Topic:</h1>
<form method="POST" action='{{urlfor "SubmitController.CreateTopic"}}'>
  {{ .xsrfdata }}
  {{template "alert.tpl" .}}

  <label for="inputName">Name</label>
  <div>
    <input placeholder="name" name="Name" value="{{index .Params " Name "}}" type="text" id="inputName" />
  </div>

  <label for="inputTitle">Title</label>
  <div>
    <input placeholder="title" name="Title" value="{{index .Params " Title "}}" type="text" id="inputTitle" />
  </div>

  <label for="inputDescription">Description</label>
  <div>
    <input placeholder="description" name="Description" value="{{index .Params " Description "}}" type="text" id="inputDescription" />
  </div>

  <input type="submit" value="Create">
</form>