<?php

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\RedirectResponse;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use Symfony\Component\Process\Process;


$app->get('/', function () use ($app) {
    return $app['twig']->render('index.html', array());
})->bind('homepage');

function getNames($s) {
  $L = split(',', $s);
  $n = count($L);
  if ($n) {
    $s = $L[0];
    for ($i = 1; $i < $n; ++$i)
      $s .= '\\\\ ' . trim($L[$i]);
    return $s;
  } else {
    return '';
  }
}

$app->post('/', function(Request $request) use ($app) {
  $context = $request->request->all();
  if (array_key_exists('name', $context))
    $context['name'] = getNames($context['name']);
  $tex = $app['twig']->render('caratula.tex', $context);
  if (array_key_exists('tex', $context))
    return new Response($tex, 200, array('Content-Type' => 'text/plain'));
  $location = __DIR__ . '/../web/tmp/';
  $tmpdir = exec('mktemp -d -p ' . $location);
  exec('cp ' . __DIR__ . '/../templates/ucsp.png ' . $tmpdir);
  $comp_pr = new Process('pdflatex', $tmpdir, array('PATH' => '/usr/bin'), $tex);
  $comp_pr->run();
  $response = null;
  if ($comp_pr->isSuccessful()) {
    $file = fopen($tmpdir . '/texput.pdf', 'r');
    $pdf = fread($file, filesize($tmpdir . '/texput.pdf'));
    fclose($file);
    $response = new Response($pdf, 200, array('Content-Type' => 'application/pdf'));
  } else {
    $response = new Response($comp_pr->getOutput(), 200, array('Content-Type' => 'text/plain'));
  }
  exec('rm -r ' . $tmpdir);
  return $response;
})->bind('caratula');


$app->error(function (\Exception $e, $code) use ($app) {
    if ($app['debug']) {
        return;
    }

    // 404.html, or 40x.html, or 4xx.html, or error.html
    $templates = array(
        'errors/'.$code.'.html',
        'errors/'.substr($code, 0, 2).'x.html',
        'errors/'.substr($code, 0, 1).'xx.html',
        'errors/default.html',
    );

    return new Response($app['twig']->resolveTemplate($templates)->render(array('code' => $code)), $code);
});
