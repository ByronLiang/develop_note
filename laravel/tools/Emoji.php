<?php

class Emoji
{
    /**
     * Emoji原形转换为String
     * @param string $content
     * @return string
     */
    public static function encode($content)
    {
        return json_decode(preg_replace_callback("/(\\\u[ed][0-9a-f]{3})/i", function ($maps) {
            return addslashes($maps[0]);
        }, json_encode($content)));
    }

    /**
     * Emoji字符串转换为原形
     * @param string $content
     * @return string
     */
    public static function decode($content)
    {
        return json_decode(preg_replace_callback('/\\\\\\\\/i', function () {
            return '\\';
        }, json_encode($content)));
    }

    /**
     * Emoji字符串清清理
     * @param string $content
     * @return string
     */
    public static function clear($content)
    {
        return preg_replace_callback('/./u', function (array $match) {
            return strlen($match[0]) >= 4 ? '' : $match[0];
        }, $content);
    }

    /**
     * Emoji字符串是否存在
     * @param string $content
     * @return Boolean
     */
    public static function isMatchEmoji($content)
    {
        $pattern = '/./u';
        $rs = preg_match_all($pattern, $content, $match);
        if($rs > 0){
            foreach($match[0] as $m){
                if(strlen($m) >= 4){
                    return true;
                }
            }            
        }

        return false;
    }
}

$content = 'ko👑🎤哈罗';
$res = Emoji::isMatchEmoji($content);

echo $res ? 'have emoji' : 'not have emoji';

if ($res) {
    $text = Emoji::clear($content);
    echo "\n".$text;
    $encodeText = Emoji::encode($content);
    echo "\n".$encodeText;
}