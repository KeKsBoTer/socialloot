<div class="auth-form">
  <h1>create topic</h1>
  <form method="POST" action='{{urlfor "APIController.CreateTopic"}}'>
    <label for="inputName">Name</label>
    <div class="reverse-order">
      <input name="name" value="" id="autofocus" type="text" id="inputName" autofocus=/>
      <label for="inputName" class="message"></label>
      <p class="description">The topic's short name<br> Only alpha characters and numerics are allowed</p> 
    </div>

    <label for="inputTitle">Title</label>
    <div class="reverse-order">
      <input name="title" value="" type="text" id="inputTitle" />
      <label for="inputTitle" class="message"></label>
      <p class="description">The topic's full name</p>
    </div>

    <label for="inputDescription">Description</label>
    <div class="reverse-order">
      <input name="description" value="" type="text" id="inputDescription" />
      <label for="inputDescription" class="message"></label>
      <p class="description">Short description for the topic</p>
    </div>

    <p class="global-message"></p>
    <input type="submit" value="Create Topic">
  </form>
</div>