<?php

namespace Caratula\Utils;

use Silex\Application;

class FileManager
{
    private $path;
    private $dir;

    public function __construct(Application $app)
    {
        $storage_config = $app['config']['storage'];
        $this->path = $storage_config['path'];
        $this->dir = $storage_config['dir'];
    }

    public function upload($file, &$url)
    {
        $rfile = "";
        do {
            $rfile = $this->easyRandom();
        } while($this->exist($rfile));

        $result = copy($file, $this->dir . $rfile . '.pdf');

        $url = $this->path . $rfile . '.pdf';
        return $result;
    }

    public function exist($file)
    {
        return file_exists($this->dir . $file . '.pdf');
    }

    public function getURL($file)
    {
        return ($this->exist($file)) ? $this->path . $file . '.pdf' : false;
    }

    public function easyRandom()
    {
        $alphabet = array(
            "bcdfghjklmnpqrstvwxyz",
            "aeiou"
        );

        $string = "";
        for($i = 0; $i < 6; ++$i)
            $string .= $alphabet[$i%2][rand(0,strlen($alphabet[$i%2])-1)];

        return $string;
    }
}