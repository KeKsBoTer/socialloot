<div class="auth-form">
  <h1>create topic:</h1>
  <form method="POST" action='{{urlfor "ApiController.CreateTopic"}}'>
    <label for="inputName">Name</label>
    <div>
      <input name="name" value="" type="text" id="inputName" autofocus/>
    </div>

    <label for="inputTitle">Title</label>
    <div>
      <input name="title" value="" type="text" id="inputTitle" />
    </div>

    <label for="inputDescription">Description</label>
    <div>
      <input name="description" value="" type="text" id="inputDescription" />
    </div>
    <p class="message"></p>

    <input type="submit" value="Create">
  </form>
</div>