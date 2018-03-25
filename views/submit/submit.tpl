<h1>submit:</h1>
<form method="POST" action='{{urlfor "ApiController.Submit"}}'>
  <label for="inputTitle">Title</label>
  <div>
    <input placeholder="title" name="Title" value="" type="text" id="inputTitle" />
  </div>

  <label for="inputContent">Text</label>
  <div>
    <textarea placeholder="content" name="Content" type="text" id="inputContent"></textarea>
  </div>

  <label for="inputTopic">Topic</label>
  <div>
    <input placeholder="topic" name="Topic" value="" type="text" id="inputTopic" />
  </div>
  <p class="message"></p>

  <input type="submit" value="Submit">
</form>