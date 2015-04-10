<?php

namespace Caratula\Utils;

class FileManager
{
    private $host;
    private $user;
    private $password;
    private $path;

    public function __construct()
    {
        $this->host = getenv('CARATULAS_FTP_HOST');
        $this->user = getenv('CARATULAS_FTP_USER');
        $this->password = getenv('CARATULAS_FTP_PASSWORD');
        $this->path = getenv('CARATULAS_FTP_PATH');

        //echo "HOST: " . getenv('CARATULAS_FTP_HOST') ." EOF";die();
    }

    public function upload($file)
    {
        $conn_id = ftp_connect($this->host);

        $login_result = ftp_login($conn_id, $this->user, $this->password);

        $rfile = "";
        do{
            $rfile = $this->easyRandom();
        }while($this->exist($rfile));

        $result = (ftp_put($conn_id, $rfile . '.pdf', $file, FTP_BINARY)) ? true: false;

        ftp_close($conn_id);

        return $result;
    }

    public function exist($file)
    {
        $conn_id = ftp_connect($this->host);

        $login_result = ftp_login($conn_id, $this->user, $this->password);

        $result = (ftp_size($conn_id, $file . '.pdf') != -1) ? true : false;

        ftp_close($conn_id);

        return $result;
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
        for($i=0;$i<6;$i++)
        {
            $string .= $alphabet[$i%2][rand(0,strlen($alphabet[$i%2])-1)];
        }

        return $string;
    }
}