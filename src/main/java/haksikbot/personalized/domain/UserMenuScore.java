package haksikbot.personalized.domain;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Getter @Setter
@Entity
@Table(name = "user_menu_scores")
public class UserMenuScore {

    @Id
    @Column(name = "id")
    private Long id;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user; // 유저

    @ManyToOne
    @JoinColumn(name = "menu_name")
    private Menu menu; // 메뉴

    @Column(name = "score")
    private int score; // 점수

    @Column(name = "count")
    private int count; // 횟수

    //==생성 메서드==//
    public static UserMenuScore createUserMenuScore(User user, Menu menu, int score) {
        UserMenuScore userMenuScore = new UserMenuScore();
        userMenuScore.setUser(user);
        userMenuScore.setMenu(menu);
        userMenuScore.setScore(score);

        return userMenuScore;
    }

    /**
     * 메뉴별 점수 업데이트 로직
     * 값이 이미 존재하는지에 대한 여부는 해당 코드에서 판단 X
     * (기존 점수 * count + 새로운 점수) / (count + 1)
     */
    public void updateMenu(int newScore) {
        this.score = (this.score * this.count + newScore) / (this.count + 1);
        this.count++;
    }

}
