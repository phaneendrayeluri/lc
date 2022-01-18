package challenges;

import java.util.HashMap;
import java.util.Map;

// 290
public class WordPatterns {

    public static void main(String[] args) {
        System.out.println(wordPattern("abba", "dog cat cat dog"));
        System.out.println(wordPattern("abba", "dog cat cat fish"));
        System.out.println(wordPattern("aaaa", "dog cat cat dog"));
        System.out.println(wordPattern("abba", "dog dog dog dog"));
        System.out.println(wordPattern("ccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccdd", "s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s s t t"));
    }

    public static boolean wordPattern(String pattern, String s) {

        String[] words = s.split(" ");
        HashMap result = new HashMap();

        if (pattern.length() != words.length) {
            System.out.println("hello");
            return false;
        }

        for (int i = 0; i < pattern.length(); i++) {

            char c = pattern.charAt(i);
            String w = words[i];

            if (!result.containsKey(c))
                result.put(c, i);

            if (!result.containsKey(w))
                result.put(w, i);

            if (!result.get(c).equals(result.get(w))){
                System.out.println(result.get(c) + " " + result.get(w) + " " + c + " " + w);
                return false;
            }
                
            
        }

        return true;
    }

}
