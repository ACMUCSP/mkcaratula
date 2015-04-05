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

function joinNames($L) {
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

function filterQuote($s) {
  $single = false;
  $double = false;
  $r = '';
  for ($i = 0; $i < strlen($s); ++$i) {
    if ($s[$i] == '"') {
      $r .= $double ? "''" : "``";
      $double = !$double;
    } elseif ($s[$i] == "'") {
      $r .= $single ? "'" : "`";
      $double = !$single;
    } else {
      $r .= $s[$i];
    }
  }
  return $r;
}

function processContext($context) {
  if (array_key_exists('name', $context)) {
    $L = explode('/', $context['name']);
    $context['name'] = joinNames($L);
    $context['number'] = count($L);
  } else {
    $context['number'] = 1;
  }
  if (array_key_exists('title', $context)) {
    $context['title'] = filterQuote($context['title']);
  }
  $categories = array(
    'inmasc' => 'El alumno declara',
    'infem' => 'La alumna declara',
    'grupal' => 'Los alumnos declaran');
  $context['cat'] =
      array_key_exists('cat', $context) &&
      array_key_exists($context['cat'], $categories) ?
          $categories[$context['cat']] : $categories['inmasc'];
  return $context;
}

$app->post('/', function(Request $request) use ($app) {
  $context = processContext($request->request->all());
  $tex = $app['twig']->render('caratula.tex', $context);
  if (($context['tex'] ?: 'false') != 'false')
      return new Response(
          $app['twig']->render('latex_gen.html', array('code' => $tex)));
  $location = __DIR__ . '/../web/tmp/';
  $tmpdir = exec('mktemp -d -p ' . $location);
  $comp_pr = new Process('pdflatex', $tmpdir, array('PATH' => '/usr/bin'), $tex);
  $comp_pr->run();
  $response = null;
  if ($comp_pr->isSuccessful()) {
    $file = fopen($tmpdir . '/texput.pdf', 'r');
    $pdf = fread($file, filesize($tmpdir . '/texput.pdf'));
    fclose($file);
    $response = new Response($pdf, 200, array('Content-Type' => 'application/pdf'));
  } else {
    $response = new Response(
        $app['twig']->render('latex_error.html', array('code' => $comp_pr->getOutput())),
        400);
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
