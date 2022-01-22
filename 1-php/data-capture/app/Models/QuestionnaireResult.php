<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Str;

class QuestionnaireResult extends Model
{
    use HasFactory;

    public $incrementing = false;

    protected $keyType = 'string';

    protected $hidden = ['id'];

    public static function boot()
    {
        parent::boot();

        static::creating(function ($result) {
            $result->id = Str::uuid()->toString();
        });
    }

    public function questionnaire()
    {
        return $this->belongsTo(Questionnaire::class);
    }
}
