package com.alanlyne.courseapi.topic;

import java.util.Arrays;
import java.util.List;
import org.springframework.stereotype.Service;

@Service
public class TopicService {

    private List<Topic> topics = Arrays.asList
            (new Topic("CS425", "Audio & Speech Processing", "Joe plays music"),
            new Topic("CS402", "Parralell & Distributed System's", "Dermot doens't look at the readMe"),
            new Topic("CS355", "Intro to Computation", "Complain to get full marks"),
            new Topic("CS424", "Design & Semantics", "Brackets"));

    public List<Topic> getAllTopics() {
        return topics;
    }
}