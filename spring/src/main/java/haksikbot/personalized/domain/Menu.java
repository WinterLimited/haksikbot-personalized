package haksikbot.personalized.domain;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Getter @Setter
@Entity
@Table(name = "menus")
public class Menu {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "menu_name")
    private String menuName;

    @Column(name = "score")
    private int score;

    @Column(name = "count")
    private int count;

    //==생성 메서드==//
    public static Menu createMenu(String menuName, int score) {
        Menu menu = new Menu();
        menu.setMenuName(menuName);
        menu.setScore(score);
        menu.setCount(1);
        return menu;
    }

    /**
     * 메뉴별 점수 업데이트 로직
     * 값이 존재하는 경우 동작함
     * score = (score * count + newScore) / (count + 1)
     */
    public void updateMenu(int newScore) {
        this.score = (this.score * this.count + newScore) / (this.count + 1);
        this.count++;
    }
}
