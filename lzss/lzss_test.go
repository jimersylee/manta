package lzss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert := assert.New(t)

	output, err := Uncompress([]byte{})
	assert.Equal("EOF", err.Error())
	assert.Nil(output)

	output, err = Uncompress([]byte{})
	assert.Equal("EOF", err.Error())
	assert.Nil(output)

	output, err = Uncompress([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal("header was `\x00\x01\x02\x03`, expected `LZSS`", err.Error())
	assert.Nil(output)

	output, err = Uncompress([]byte{'L', 'Z', 'S', 'S', 255, 255, 0x83, 0x01, 0xb8})
	assert.Equal("EOF", err.Error())
	assert.Nil(output)

	output, err = Uncompress([]byte(sampleInput))
	assert.Nil(err)
	expect := uncompressReference(sampleInput)
	assert.Equal(output, expect)
}

/*






























*/

var (
	sampleInput = "LZSS\240\027\000\000\000\203\001\270\001\000!\000\"\000\322P\330\000C\000\004\001\000\3052\240\010\000\000\360\007B*\001\0037\000`1\000B \000\004@D\032\n\033`\014\000\200$\240\214\006\024\001\000\020\000\376@E\001\002\340\006\000\004L\006\000B\204\000\010H\303\000\377\000\014\001\220\005\024\313\200\200\"\000\000\300\037(\001\003\020\334\000\200\315\000C\020\000\001\000i\370\037\2001\000\302\200\0002\032P\004\000\000\370\003B\005\001\003\033\0000\032\000B\020\000\002 \036\r\355\t0\004\000@\034P,\003\212\000\000\010\000\177\240\001\003p\003\000V\002\003\000BB\000\304\243\241=\000\001\306\000\310\003\312h@@\021\000\000\340\017\324\001\003n\000\000\300l\000n\000@\010\000\200p4$\025\300\020\000\000\211@\261\014(\002\000\000\010\374\201r\001\003\r\000\330\r\001\000B\010\001\020\216\206\244\002\000\030\003 \022(\243\001E`\000\000\200?P\001\003\017r\303\001\017\305\242\321\260\277\000#\000\264d\002\017\307\252\001\003\017\3029\017\305\0004\032\366\027`\010\200PM\017\3105\001\002\017r,\006\017\326F\200C\373\003\014\001\020\013\017\330\023\001\003\017\202\305\304\017\345\301hhR\177\017\342r\201\017\347\305\001\003\033\020\000\260\030\031\017\365\026\rA\340\0130\002@0\017\370\001\003\017\242\004\0263\020\006\242!h\001\2068\000H\006\020\010\001\003\020\002bh\241\020\025P4\014\024\020\022\331\020\030BR\001\003\r\000XL\020&\212t\206\201\020#\034\020(\001\003\017\322\213\010\261\001\230 \004\321\360\t#8\000\244\003\037\370\000\362\037\3437\000\0023\020#$\032>a\010\200Bx\020\030%\032>\341\020\003\007\321\020\005DC\366\020\003\020\020\010\001\003\205\020\003\344\020\005\201h\310~\020\0024\022\202\020\007\205\001\003\037\362\031\030D\200\031\020\003\016\ri\037\362F'\017\370\000\362\017\202&\023\017\345\204\241Da\273\017\343\t\212e\017\345\224\023\001\003\017\342dd\017\345(44R'\017\3429A\017\347*\001\003\r\020\000\230\314\014\017\345\204\206\340\200\003\030\002 )(\226\017\346\023\001\003\017\342\223\241 %b\3200\322S/\362d\005\017\367j\001\003\020\002\20425 5\010\032\026\016 2h\264\240X\020\026\t\001\0020\003\306\241 5\210g\300\3650\023\027\020'&\250\001\003\017\302\311\3340&\rH\000?\200!\000\022\203b\031\0270&\001\003\020#\034 5\242\001m\272\017\0202f\0208\001\003\020C\223\020F\3504\240k\020C\r\020H\001\003\020B\020f`\000f\020C\230\0064\320\303\020\000\271\0208\232\000\36202\242L\020&\324\200\251\020#9\020(\023\001\003\017\302\233\2210V\032P\\\235\020#\0070X\001\003\020\"33\020%PT\003\212\013@2\360 HU\023\001\002\017\302l\206\020&l@g] S\037\020(\001\003@C\324\020%\221H\r\350l R\362\203 W\225\023\001\0030b\031\033\020%\272\001\323:\013\020\"\202\020(\001\003\017\3026s\221\020&7 )\020\"\310\020\020(\013\001\003\020#p u\370\006\244\027i r)B\0207\372\001\003\020C\016D\300\014\020C\340\200J\0202G'\0208\000\362\017\302\243\201PV\034pr9\020\"$\t0X\001\003\020\"4!Pf\210\003z\006 R,\241M W\211\001\002\017\302\214F\020&qt@\326 S&\020(\001\003\017\302\321\010\314\000\314\020#A\016H\202\320!\000\362\204 GE\000\362Pc\241Pv\312\0019\005\020\002\242\020\010\023\001\003\017\242FS\020\0069`\372]\020\003\024\020\010\001\003\020\002hP\2268\250\007L6\020\002\251\020\010:\001\003\020\r\000\030\315@v\350\200\240\272\002\020\022W\020\030\001\003\020#\301\020&\350\035\220\267\020#\013\020(\001\003\020#\241P\306\250\003\216\023\020\"l\020(\026\251\001\002\017\302\254P\306\210u@r\376 R\220.\020(\001\003\020\"\325\221P\307\016\010\0340\202\362\205\0207\026\305\001\003 R\032P\306\332\001\275=`\303\302\020H\001\003`\303P\306\204;\350\340\354\000\020B\030\020H\001\003\020BBjP\306x\007\214= SC-\020Gz\001\003P\302M\020F\360\200t\211\004\020Bg\020H\001\003\017\342\253\242\261\020F\036\220\336\020C\r\020H\023\001\003\020B57\020E\310\003\002\352\033\020B\254\020H\311\001\002\020DQ\005\320\210y@}\020C6\020H\001\003C\020CQ\006A\017(\177\020C\206' \210\001\003 \222\033\030\020E\352\001\272\013@\323\342\020H\001\003\017\342fQ'\350=`\265\020C\034\020H\001\003q#QQ&\270\007\344a\023\251\020H\272\013\001\003Q\"\315@\366\370\200\334\005]\020Bw\020H\001\003\017\342\263Q&\"\350\037p\375\020C\017\020H\001\003\020B\2426Q&\350\003\2660\303\354\020H\026\351\001\002\017\342\314Q&\010}\300\000\326\000\214\001\020>\224\321'aV\001\003\017\342\331\3340\305\261\017@\010\207!\000\342\207\0207\265\013\000\362\020#\034a5\366\001\3410\340\006@\376PF\020\026\000\362\020\003\010\223\003P\020\003\004?\240\n\240F\000\010`\000\020\007\024\001\002!\017\362n`\000j\017\363\340\007\000T\301\020\000\t\014@\031\032\r\017\365\342\001\002@\222M\014@\"\016\017\343\375\200\241\017\343\203\001\207\017\370\001\023\017\222\273\221\001\310P\364\000\0370\264\000c\000\20408\000e4Q\006\001\023\020\00273\004\0009\020\003\370\003\346\n`\320\004\200\024\006aH\371\001\022\017\262 \354\206\006 \007\020\023\177\300\344\\\001Q\022\303\000\020(\001\023\020\"\020\335\324\000\344\020#\361\017\210\200(\200\021\000r\030\200\020G\206\365\001\0230\322\033\033\200\034\020S\310\376\001\021A#\020\003\020x\001\023!\020\202vs\003\220\020\204`\340r\355 \222\210b\020\231\001\023\020\243p\004\000r\020\243\000\014\274\035\300h\030\000Y\020\271\002\001\023\020\303\016\321\020\305\212\202\3641\003\215\020\311\001\023\205\222\023\201\020\305BQ\220\236\020\3024\3041\020\310J\001\023\020\30281\241\020\305\270\016\336\032 \322@\020\311&\271\001\022\020\302\014G\020\306\327\301\300[\003\214\001\220\310\020\311\001\023\211\020\302\341\314\020\306:\310n1B\2342\031\020\311\001\023!\"\034\032\020\305\320^\007\331\r!2(\020\311\001\023\211\020\302\206S\020\306\354\340~\020\303Ne\020\311\001\023\020\302pl\020\305\200\250\035\334\017\020\302\271\020\311\202\001\023\211R\"\316\r\020\305\261\203\237R#.\231\020\311\001\023\020\303\301\020\305\"v\244\360S\020\302D3\020\310*\001\023E\020\3039\020\305\310\016\272\202\323p-\020\311\311\001\022\243#\007\020\306\345\301\200\370\000\214\000\220\316\000\223(\023\001\023\020\302\345\304\020\305\241<\030B\037\020\302\342\031\200\242\020\306\245\023\001\023!\222\034\031\020\305\224\007\343\302\003\020\302>\003PH\020\306\001\023\t\020\302\2263\020\305\304\362\340>j\00122h\020\310\324\001\023\020\302r\002h\020\305`\036\244\026\300\010h\000\t\r\243xb\001\023R\242N\201\020\306\315\203\231\004\030\001\211\302'\020\310\001\023\020\302\313\261\020\305\302y\244\260\334\263\3034\000Bw\312\001\023E\263\3237\020\305X\017\3562c\220-\020\311Y\001\022\020\304\007\020\305\353\301\202]2b\220\322\000\024\315\020\306\013\001\023\020\303\344\020\305a=\270+\241\203\222b\032\200B\203\246e\001\023\t\264\023\030\030\020\325\266\007\331\t9\264\"N\003S\030\001\023\2643\003\023\241\020\345D\367\340\246\020\343j\020\350&T\001\023\264S`d2\226\037\254j\036\020\362I\020\371\002\001\023\020\362\014B\314Cv\344\203\320\002\021\002\252\007\021\t\001\023\264\223\201\241\0018\002\001\324\267\203\000d5\000E5\000\000&\240\210930\010!\324\344k\222\260v\325405\004\000G\021\223@\032\376\007`\000\020\200\260\006\240\250\006\300\000\004\0241g\006\006AA\221\0013Q\022\022d\343,\006T\242\002\010d\363GC{\002\014\002\000\220\326\000\024\325\000\230\200@\"\346\314\300 \250\0013\254\020J\302r\002\266\203\300\334\000\004\034\001\022\222\341hH*\200\000A\000\342\032\200\242\032\000\000\023P\304\234\031\030\004\345!\0013eI\250]\023\005\034\200\002#\023\003\030\r\355\0170\010\000@^\003PT\003`\002@\212\2303\003\203\240\0013\260\250.\t\311\307\264\003\247\202\004f\243\000\242a\240\000\006\001\010l\000\000\212j\000L@\021s\020f`\020\224\0013\026&\341\202\2775\243bb`\000\216\023\343\000@4d?\300 \000\211\000\r@Q\r\200\t(b \316\014\014\202B\0013\312$,\004>\310\324\250\302\017\024S\207\206\0024\310\322\262\001(\252\0010\200\001E\314\231\201AP\001\"DX\232%\223\213\211\221\024\205\302H\320\260\3356rd6\024\212\312!\0013k\223\360\275\024\20413\241\024\205\024\032\232\023\251\242\320\024\213B\025\0013q\0226\017\024\203&\203\2522\024\203\010BC\360\001HB\n\332\024\212(\0013\254NB\355\205\333%\324\024\2051h\230)'#\212\033\024\2125\0013\345I(\252\343\020\260\230\030\033\024\205\004\r\013\032\0078rn\024\213\0013\260>\t\032\023(\004\023\253\342\024\203\3043\340@z\000\306\000\010n\024\212\324\241\0013\026(\341\273\024\205p\024\205P\200\006\244\037JB\311\024\213\202\261\0013\n%t\335e\255\002\021\024\203\320\321\200\266\007\024\222\272\024\233\0013\220X\242\204\361)T\221\201\024\305PB\032\320\265\024\3037\024\312J!\0013k\224\020\374\024\30421\303~\342\024\302L\003\232a\276b\024\273BM\001\"\221\022f\022\024\243F\243\256\262\024\244j\300\324\024\243\336\024\253!\0013\254RBl)4\310\314\204\000\024\024\243Q\r(.)3\205\024\253U\0013eJX\204\360TB\032\024\225\262\001\235\r)\"~C\024\233\0013\260\\\t-^S\026\006#\260r\024\223D7`z\001\025\024\222p\024\232T\0013\226,!\022\257\024\224dl\024\225\360\006$R%\024\222\t\016\024\232\362\0013\232\210%\034\016\024\223\214\314\r@\302\321\024\222\337\200\364\362\023\302\024\233\0013\020X\264\204|\024\204\301\001\030\241\024\203\002\034P\t\024r8\024zB\n\001\"\253\226p^\024e9!)5\204\003.\007)2\020\007E\024z\205\0013\331\022V\263\003,Ff\263\022\024\204q@\317\222r\220\206\342\024\213\0013\254[\302\345\024\204D\314\304)6\016\310\032)3\034C);\0013\205KHx)4\031F\031c2\024\222\310\001I)\"\216C\024\213\001\"\260r\ta\024t3\243\264\242\024t9 \247)\023r\024{!\0013\226.a\367\024tfh\241\024u0\007L?\024rI\024{\"2\0013\332%\214c\004\314LB\r)\025\347\200\311\006\024r\312\203\024{\0013X\274\204\237\006\024\203\204\231\261)\025\002\035\020T)\023\2069\024\233\0013\253\227p\367\024\244\20437\024\245\244\003\362\026\024\242\2120\024\253\245\0013\371\022J\246\263\r\024\242\007g\262\024\242u\300q\002\r\024\242\346\024\253\0013\254_B-\205\024\224\344\024\225\261\016\310\177\024\222\n\342\024\233\265\0013\005L\010m\211Rd\032\030)5\330\001\201vS\206\236\024\253\0013\260\202\t\005RdFC\271\022\024\264;\2407)3t#\024\273\0013\2260a\271\224bh\242d\024\265p\007\234\230\363\211\024\273Br\0013\032&\354\034f\304\315\242\014)5\357\200\261=\323\322\024\273\241\0013X\311\204d\024\243\241g&\320\002\0360\221\024\243:\024\253\0013PK\2310\363\024\2444g&\304\250\003\322\033\024\242P\024\253\305\0013(-\023\346R\224\206g&\010y4@`\232\343\352\024\253\0013,f\024\302\346R\224\320g&1\017\250\032o)S\035)[\0013\325L(\2069\024\265g&\350\001\345\017\024\262\206\256\024\273\0013\260\234\t\263\315\303\204\026Cg&D=`\341\024\263\206v\024\273\0013\3263\341\360\024\264\302jg&\260\007\254\026R\223\024\273B\262\0013\202&T$R\224M\242\014)u\367\200\234=\363\332\024\273\241\0013X\321\204\214gd\251gf\320\002\037\220\273\024\303;\024\313\0013PK\232p\177\024\3045gf\344\250\003\256\037\024\302p\024\313\345\0013\350M\023*>$\246gg\317\223{\362\206\356\024\313\0013,lBX\216\203\340\305\324\324\000\014)\204\3202R\242C\024\273\321c\225Mx?>\004\032\241gF\372\001C\013R\242\276\024\253!\0013\260\264\t\365\301\323\026S7gF\301\344{\322x\024\252\302\004\3266\024\341\370\217\304jgFP\024\244R'{\322\t\017\024\252R\0013\342\030&\364-\024\245gF\256\203\267`\006\030\004 \342\024\253\0013X(\335\204\205gD\261gF\342uT\220\335\220\243<\024\252\352\0013\313H\233p\273\024\24461\024\245\300T\016\356\220\243\220\024\253\301\0013}D\023\332>$\306F\006>$\210\320\330\301O\001\024\242\362\220\253\0013P,p\302\032\024\223\330gV!\250;\350:)c\036\024\252%\0013H\025N\370gT\033\032\n\202\251\252\001\000\202\033\000\202\271\000\203\034\000\202\000\311\001\260\230\033\030\000\213\324\271\211\000\203\031\000\202\231\000\203\002\302\252\271\002\302\230\002\303\271\002\302\230\002\303R\271\002\302\030\034\002\302\301\002\302\030J\034\002\302\301\002\302\030\034\002\302\301U\005\223\034\002\302\301\005\223\034\002\302\301\333\005\223\002\303\311\005\223\002\303\311\005\223\002\3036\311\005\223\002\303\311\005\223\002\303\311\311@\0010\031\030\030\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000"
)
