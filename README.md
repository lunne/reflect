# reflect
Reflect on your work each day by doing a quick check-in in the morning and a reflection and report of the day in the afternoon.

Goals usually stretch over days, weeks or years. Something that should be done each day and reflected upon. This is done by setting up recurring goals that will be visible and easy to select in the check-in.

At the end of each week (or two weeks) a report from the system can be created to do a deeper analys of the progress. Are you completing the goals you setup or are you failing to even work on them? This will help you create new goals to help you correct and improve your performance.


The user checks in in the morning by doing a simple cli command:
`reflect checkin`

This will start up the user workflow that:
  1) Creates goals from any recurring goals the user have created
  2) prompts the user to set new goals

```

   _ \          _|  |              |
  |   |   _ \  |    |   _ \   __|  __|
  __ <    __/  __|  |   __/  (     |
 _| \_\ \___| _|   _| \___| \___| \__|


Previously, you’ve described the following goal to improve/maintain good habits at work (recurring goals).
Do you want to add them for today as well?
    [ ] Finnish one code review before lunch
    [ ] Focus on work and silence any social media (facebook,youtube,messenger) notifications
    [ ] Medidate for 10 minutes

Please answer the following question after planning your workday and before starting with your work.
What do you plan to achieve today? (This could include tasks and other things you want to make progress on at work)
    [n] Create new Goal
    [c] No more goals for today
         __
 _(\    |@@|   Have a conscious and productive day!
(__/\__ \--/ __
   \___|----|  |   __
       \ }{ /\ )_ / _\
       /\__/\ \__O (__
      (--/\--)    \__/
      _)(  )(_
     `---''---`
```


At the end of each day the user can write the command `reflect check-out`

---
```

   _ \          _|  |              |
  |   |   _ \  |    |   _ \   __|  __|
  __ <    __/  __|  |   __/  (     |
 _| \_\ \___| _|   _| \___| \___| \__|


Please rate whether you achieved what you set out to do this morning for your following items [(C)completed, (P)rogress, (N)o progress]:
    [C, P, N] Finnish one code review before lunch
    [C, P, N] Focus on work and silence any social media (facebook,youtube,messenger) notifications
    [C, P, N] Medidate for 10 minutes

---
What was the best work-related thing that happened to you today at work?
[Text]

---
Was there anything else that you achieved but didn’t plan for in the morning?
[Create and complete goal]

---
Was there something that made it difficult to achieve what you planned to do? Previously, you’ve described the following goal to improve or maintain good work habits:
    - Finnish one code review before lunch
    - Focus on work and silence any social media (facebook,youtube,messenger) notifications
    - Medidate for 10 minutes

---
Did you actively try to achieve the goal?
    [yes, no]

[show if 'yes' ] 
How did you carry out your strategy?
[text]

Did it positively influence your workday?;
[text]

[show if ’no’] Why not?
[text]
---

Is there a goal that you would like to set for yourself that could help you to improve or maintain good work habits? You can revise your existing goal, set a new goal, or keep the same goal?

[Show information about SMART goals]
    Try to follow the SMART goals principle, i.e., think about a goal that is specific, realistic to achieve and matters to you.

    It might be helpful if you also think about when you want to achieve that goal (time-bound) and how you could measure your progress towards it. [copy previous goal into textbox if available]

---

If 'yes' (new goal):
What’s the first step you will take towards reaching this goal?
[text]

---
Did any other goals come to mind today that you would like to mention?
[text]

---
How did you approach assessing your achievements and the progress you’ve made today?

[text]

         __
 _(\    |@@|   Thank you and see you tomorrow!
(__/\__ \--/ __
   \___|----|  |   __
       \ }{ /\ )_ / _\
       /\__/\ \__O (__
      (--/\--)    \__/
      _)(  )(_
     `---''---`
```


## Reference
https://ieeexplore.ieee.org/document/8823032