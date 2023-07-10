<?php
require_once 'config.php';

if ($_SERVER['REQUEST_METHOD'] === 'GET') {
  // Buscar todas as tarefas do banco de dados
  $query = "SELECT * FROM tasks";
  $result = mysqli_query($conn, $query);

  $output = '';

  if (mysqli_num_rows($result) > 0) {
    while ($row = mysqli_fetch_assoc($result)) {
      $output .= '<li class="taskItem">' . $row['task_name'] . ' <button class="editTaskButton" data-id="' . $row['task_id'] . '">Editar</button> <button class="deleteTaskButton" data-id="' . $row['task_id'] . '">Excluir</button></li>';
    }
  } else {
    $output = '<li>Nenhuma tarefa encontrada</li>';
  }

  echo $output;
} elseif ($_SERVER['REQUEST_METHOD'] === 'POST') {
  // Criar uma nova tarefa
  $taskName = $_POST['taskName'];

  $query = "INSERT INTO tasks (task_name) VALUES ('$taskName')";
  $result = mysqli_query($conn, $query);

  if ($result) {
    echo 'success';
  } else {
    echo 'error';
  }
} elseif ($_SERVER['REQUEST_METHOD'] === 'PUT') {
  // Atualizar uma tarefa existente
  parse_str(file_get_contents("php://input"), $putVars);
  $taskId = $putVars['taskId'];
  $taskName = $putVars['taskName'];

  $query = "UPDATE tasks SET task_name='$taskName' WHERE task_id='$taskId'";
  $result = mysqli_query($conn, $query);

  if ($result) {
    echo 'success';
  } else {
    echo 'error';
  }
} elseif ($_SERVER['REQUEST_METHOD'] === 'DELETE') {
  // Excluir uma tarefa
  parse_str(file_get_contents("php://input"), $deleteVars);
  $taskId = $deleteVars['taskId'];

  $query = "DELETE FROM tasks WHERE task_id='$taskId'";
  $result = mysqli_query($conn, $query);

  if ($result) {
    echo 'success';
  } else {
    echo 'error';
  }
}

mysqli_close($conn);
?>
