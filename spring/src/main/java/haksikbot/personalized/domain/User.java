package haksikbot.personalized.domain;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Getter @Setter
@Entity
@Table(name = "user")
public class User {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "user_id")
    private Long id;

    // 이름
    // private String name;

    // 학번
    // private String studentId;

    // 비밀번호
    // private String password;
}
