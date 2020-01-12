<?php

require_once 'HashMap.php';

class Test 
{
    protected $wordTree;

    public function __construct() {
        $this->wordTree = new HashMap();
    }

    public function yieldToReadFile($filepath) {
        // $fp = fopen($filepath, 'r');
        // while (! feof($fp)) {
        //     yield fgets($fp);
        // }
        // fclose($fp);
        return ['apple', 'awdy', 'awow'];
    }

    public function getWords($filepath) {
        foreach ($this->yieldToReadFile($filepath) as $word) {
            echo trim($word) . "\n";
            $this->buildWordToTree($word);
        }
        var_dump($this->wordTree->size());
        var_dump($this->wordTree->keys());
        var_dump($this->wordTree->values());
    }

    protected function buildWordToTree($word = '')
    {
        if (! $word) {
            return;
        }
        $tree = $this->wordTree;
        $wordLength = mb_strlen($word, 'utf-8');
        for ($i = 0; $i < $wordLength; $i++) {
            $keyChar = mb_substr($word, $i, 1, 'utf-8');
            // 获取子节点树结构
            $tempTree = $tree->get($keyChar);
            if ($tempTree) {
                $tree = $tempTree;
            } else {
                // 设置标志位
                $newTree = new HashMap();
                $newTree->put('ending', false);
                // 添加到集合
                $tree->put($keyChar, $newTree);
                $tree = $newTree;
            }
            // 到达最后一个节点
            if ($i == $wordLength - 1) {
                $tree->put('ending', true);
            }
        }

        return;
    }
}
$test = new Test();
$test->getWords('words.txt');