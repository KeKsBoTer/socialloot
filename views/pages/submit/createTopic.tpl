<h1>Create Topic:</h1>
<form method="POST" action='{{urlfor "ApiController.CreateTopic"}}'>
  <label for="inputName">Name</label>
  <div>
    <input placeholder="name" name="name" value="" type="text" id="inputName" />
  </div>

  <label for="inputTitle">Title</label>
  <div>
    <input placeholder="title" name="title" value="" type="text" id="inputTitle" />
  </div>

  <label for="inputDescription">Description</label>
  <div>
    <input placeholder="description" name="description" value="" type="text" id="inputDescription" />
  </div>
  <p class="message"></p>

  <input type="submit" value="Create">
</form>