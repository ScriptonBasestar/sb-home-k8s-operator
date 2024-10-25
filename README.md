# SB Home K8s Operator

홈서버용 오퍼레이터

helm을 안쓴 이유
- helm템플릿이 구리다
- 희한하게 만들기도 어렵고 쓰기도 어려운데 편리하진 않다

## 동작 환경
- Debian Bookworm k3s인프라에서 동작 확인
- 홈서버 1대짜리로 가정
- 안쓰는 컴퓨터로 돌릴만한 수준
- Hostpath를 적극 활용
- 각각의 operator, helm, image를 활용(직접 만들지 않음)

## 기능(투두)
- [ ] scripton.cloud와 적극 연동 (백업)
- [ ] Dynamic DNS
- [ ] 백업 to scripton.cloud/s3
- [ ] 보안위협 알람 scripton.cloud channel/app
- [ ] 홈 http 프록시 서버(squid,..)
- [ ] 홈 DNS 서버 (??? bind9?,..)
- [ ] 홈 vpn 서버 (wireguard, openvpn...)
- [ ] openwrt 관리 (?? ssh open해서??)
- [ ] 로컬 git 서버 (gitlab, gitea, gogs)
- [ ] git repository listen(gitops)
- [ ] vault, consul
- [ ] wordpress
- [ ] bookstack
- [ ] mariadb 오퍼레이터 이용?
- [ ] postgresql 오퍼레이터 이용?
- [ ] 나스기능 (nextcloud)
- [ ] 사진 뷰어
- [ ] 중복파일 관리
- [ ] proxynd (개발용 프록시... 미러,리포도 추가되면 추가됨)
- [ ] 매니저 dashboard
- [ ] 개발용 미들웨어? (?? mysql, mariadb, postgresql, neo4j, rabbitmq, kafka, )
- [ ] traefik
- [ ] coredns
- [ ] cert-manager
- [ ] prometheus
- [ ] harbor
- [ ] argocd
- [ ] airflow
