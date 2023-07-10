$(document).ready(function() {
    loadTasks();
  
    $('#taskForm').submit(function(event) {
      event.preventDefault();
      createTask($('#newTaskInput').val());
      $('#newTaskInput').val('');
    });
  
    $(document).on('click', '.editTaskButton', function() {
      var taskId = $(this).data('id');
      var newTaskName = prompt('Digite o novo nome da tarefa:');
      if (newTaskName) {
        updateTask(taskId, newTaskName);
      }
    });
  
    $(document).on('click', '.deleteTaskButton', function() {
      var taskId = $(this).data('id');
      deleteTask(taskId);
    });
  });
  
  function loadTasks() {
    $.ajax({
      url: 'assets/includes/tasks.php',
      type: 'GET',
      success: function(response) {
        $('#taskList').html(response);
      }
    });
  }
  
  function createTask(taskName) {
    $.ajax({
      url: 'assets/includes/tasks.php',
      type: 'POST',
      data: { taskName: taskName },
      success: function(response) {
        loadTasks();
      }
    });
  }
  
  function updateTask(taskId, taskName) {
    $.ajax({
      url: 'assets/includes/tasks.php',
      type: 'PUT',
      data: { taskId: taskId, taskName: taskName },
      success: function(response) {
        loadTasks();
      }
    });
  }
  
  function deleteTask(taskId) {
    $.ajax({
      url: 'assets/includes/tasks.php',
      type: 'DELETE',
      data: { taskId: taskId },
      success: function(response) {
        loadTasks();
      }
    });
  }
  