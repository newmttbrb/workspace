#include "gmock/gmock.h"    
#include "Soundex.h"

using namespace testing;

class SoundexEncoding: public Test {
public: 
	Soundex soundex;
};

TEST_F(SoundexEncoding, RetainsSoleLetterOfOneLetterWord) { 
   ASSERT_THAT(soundex.encode("A"), Eq("A000"));
}

TEST_F(SoundexEncoding, PadsWithZerosToEnsureThreeDigits) { 
   ASSERT_THAT(soundex.encode("I"), Eq("I000"));
}

TEST_F(SoundexEncoding, replacesConsonantsWithAppropriateDigits) { 
   EXPECT_THAT(soundex.encode("Ab"), Eq("A100"));
   EXPECT_THAT(soundex.encode("Ac"), Eq("A200"));
   EXPECT_THAT(soundex.encode("Ad"), Eq("A300"));
   EXPECT_THAT(soundex.encode("Ax"), Eq("A200"));
}

TEST_F(SoundexEncoding, ignoresNonAlphabetics) { 
   ASSERT_THAT(soundex.encode("A#"), Eq("A000"));
}

TEST_F(SoundexEncoding, replacesMultipleCOnsonantsWithDigits) { 
   ASSERT_THAT(soundex.encode("Acdl"), Eq("A234"));
}

TEST_F(SoundexEncoding, limitsLengthToFourCharacters) { 
   ASSERT_THAT(soundex.encode("Dcdlb").length(), Eq(4u));
}

TEST_F(SoundexEncoding, ignoresVowelLikeLetters) { 
   ASSERT_THAT(soundex.encode("BaAeEiIoOuUhHyYcdl"), StartsWith("B234"));
}

TEST_F(SoundexEncoding, combinesDuplicateEncodings) { 
   ASSERT_THAT(soundex.encodedDigit('b'), Eq(soundex.encodedDigit('f')));
   ASSERT_THAT(soundex.encodedDigit('c'), Eq(soundex.encodedDigit('g')));
   ASSERT_THAT(soundex.encodedDigit('d'), Eq(soundex.encodedDigit('t')));

   ASSERT_THAT(soundex.encode("Abfcgdt"), Eq("A123"));
}

TEST_F(SoundexEncoding, upercasesFirstLetter) { 
   ASSERT_THAT(soundex.encode("abcd"), StartsWith("A"));
}

TEST_F(SoundexEncoding, ignorescCaseWhenEncodingConstants) { 
   ASSERT_THAT(soundex.encode("BCDL"), Eq(soundex.encode("bcdl")));
}

TEST_F(SoundexEncoding, combinesDuplicateCodesWhen2ndLetterDuplicates1st) { 
   ASSERT_THAT(soundex.encode("Bbcd"), Eq("B230"));
}

TEST_F(SoundexEncoding, doesNotCombineDuplicateEncodingsSeparatedByVowels) { 
   ASSERT_THAT(soundex.encode("Jbob"), Eq("J110"));
}

