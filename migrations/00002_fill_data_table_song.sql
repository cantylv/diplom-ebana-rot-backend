-- +goose Up
-- +goose StatementBegin
INSERT INTO song (name, release_date, text, link)
VALUES
    (
        'Shape of You',
        TO_DATE('01-06-2017', 'DD-MM-YYYY'),
        '{
            "couplets": [
                "The club isn''t the best place to find a lover\nSo the bar is where I go\nMe and my friends at the table doing shots\nDrinking fast and then we talk slow\nCome over and start up a conversation with just me\nAnd trust me I''ll give it a chance now\nTake my hand, stop, put Van the Man on the jukebox\nAnd then we start to dance, and now I''m singing like\n",
                "Girl, you know I want your love\nYour love was handmade for somebody like me\nCome on now, follow my lead\nI may be crazy, don''t mind me\nSay, boy, let''s not talk too much\nGrab on my waist and put that body on me\nCome on now, follow my lead\nCome, come on now, follow my lead\n",
                "I''m in love with the shape of you\nWe push and pull like a magnet do\nAlthough my heart is falling too\nI''m in love with your body\nAnd last night you were in my room\nAnd now my bedsheets smell like you\nEvery day discovering something brand new\nI''m in love with your body\n(Oh-I-oh-I-oh-I-oh-I)\nI''m in love with your body\n(Oh-I-oh-I-oh-I-oh-I)\nI''m in love with your body\n(Oh-I-oh-I-oh-I-oh-I)\nI''m in love with your body\nEvery day discovering something brand new\nI''m in love with the shape of you\n",
                "One week in we let the story begin\nWe''re going out on our first date\nYou and me are thrifty, so go all you can eat\nFill up your bag and I fill up a plate\nWe talk for hours and hours about the sweet and the sour\nAnd how your family is doing okay\nAnd leave and get in a taxi, then kiss in the backseat\nTell the driver make the radio play, and I''m singing like\n",
                "Girl, you know I want your love\nYour love was handmade for somebody like me\nCome on now, follow my lead\nI may be crazy, don''t mind me\nSay, boy, let''s not talk too much\nGrab on my waist and put that body on me\nCome on now, follow my lead\nCome, come on now, follow my lead\n",
                "I''m in love with the shape of you\nWe push and pull like a magnet do\nAlthough my heart is falling too\nI''m in love with your body\nAnd last night you were in my room\nAnd now my bedsheets smell like you\nEvery day discovering something brand new\nI''m in love with your body\n(Oh-I-oh-I-oh-I-oh-I)\nI''m in love with your body\n(Oh-I-oh-I-oh-I-oh-I)\nI''m in love with your body\n(Oh-I-oh-I-oh-I-oh-I)\nI''m in love with your body\nEvery day discovering something brand new\nI''m in love with the shape of you\n",
                "Come on, be my baby, come on\nCome on, be my baby, come on\nCome on, be my baby, come on\nCome on, be my baby, come on\nCome on, be my baby, come on\nCome on, be my baby, come on\nCome on, be my baby, come on\nCome on, be my baby, come on\n",
                "I''m in love with the shape of you\nWe push and pull like a magnet do\nAlthough my heart is falling too\nI''m in love with your body\nAnd last night you were in my room\nAnd now my bedsheets smell like you\nEvery day discovering something brand new\nI''m in love with your body\nCome on, be my baby, come on\nCome on (I''m in love with your body), be my baby, come on\nCome on, be my baby, come on\nCome on (I''m in love with your body), be my baby, come on\nCome on, be my baby, come on\nCome on (I''m in love with your body), be my baby, come on\nEvery day discovering something brand new\nI''m in love with the shape of you\n"
            ]
        }',
        'https://www.youtube.com/watch?v=JGwWNGJdvx8'
    ),
    (
        'Blinding Lights',
        TO_DATE('29-11-2019', 'DD-MM-YYYY'),
        '{
            "couplets": [
                "I''ve been tryna call\nI''ve been on my own for long enough\nMaybe you can show me how to love, maybe\nI''m going through withdrawals\nYou don''t even have to do too much\nYou can turn me on with just a touch, baby\n",
                "I look around and\nSin City''s cold and empty (oh)\nNo one''s around to judge me (oh)\nI can''t see clearly when you''re gone\n",
                "I said, ooh, I''m blinded by the lights\nNo, I can''t sleep until I feel your touch\nI said, ooh, I''m drowning in the night\nOh, when I''m like this, you''re the one I trust\n(Hey, hey, hey)\n",
                "I''m running out of time\n''Cause I can see the sun light up the sky\nSo I hit the road in overdrive, baby, oh\n",
                "The city''s cold and empty (oh)\nNo one''s around to judge me (oh)\nI can''t see clearly when you''re gone\n",
                "I said, ooh, I''m blinded by the lights\nNo, I can''t sleep until I feel your touch\nI said, ooh, I''m drowning in the night\nOh, when I''m like this, you''re the one I trust\n",
                "I''m just walking by to let you know (by to let you know)\nI can never say it on the phone (say it on the phone)\nWill never let you go this time (ooh)\nI said, ooh, I''m blinded by the lights\nNo, I can''t sleep until I feel your touch\n(Hey, hey, hey)\n(Hey, hey, hey)\n",
                "I said, ooh, I''m blinded by the lights\nNo, I can''t sleep until I feel your touch\n"
            ]
        }',
        'https://www.youtube.com/watch?v=4NRXx6U8ABQ'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE song;
-- +goose StatementEnd
