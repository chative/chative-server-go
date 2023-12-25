package sms

// type B2M struct {
// 	appID, appKey string
// 	redisCmd      redis.Cmdable
// }

// func NewB2M(cfg Config) *B2M {
// 	return &B2M{
// 		appID:    cfg.B2M.AppID,
// 		appKey:   cfg.B2M.AppKey,
// 		redisCmd: rediscluster.GetRedis(),
// 	}
// }

// func (b *B2M) VerifyCode(phoneNumber, code string) error {
// 	key := "vCode:sms:" + phoneNumber
// 	val, err := b.redisCmd.Get(context.Background(), key).Result()
// 	if err != nil {
// 		return err
// 	}
// 	if val != code {
// 		return errors.New("invalid code")
// 	}
// 	return b.redisCmd.Del(context.Background(), key).Err()
// }

// func (b *B2M) SendVerificationCode(phoneNumber string) error {
// 	// 生成验证码
// 	vCode := fmt.Sprintf("%06d", rand.Intn(1000000))
// 	// 存放到redis
// 	err := b.redisCmd.SetEX(context.Background(), "vCode:sms:"+phoneNumber, vCode, time.Minute*10).Err()
// 	if err != nil {
// 		return err
// 	}
// 	// 发送验证码
// 	req, err := http.NewRequest("GET", "http://www.btom.cn:8080/simpleinter/sendSMS", nil)
// 	if err != nil {
// 		return err
// 	}
// 	q := req.URL.Query()
// 	q.Add("appId", b.appID)
// 	tNow := time.Now()
// 	tNow = tNow.UTC()
// 	ts := tNow.Format("20060102150405")
// 	q.Add("timestamp", ts)
// 	h := md5.New()
// 	h.Write([]byte(b.appID + b.appKey + ts))
// 	d := h.Sum(nil)
// 	q.Add("sign", hex.EncodeToString(d))
// 	q.Add("mobiles", phoneNumber)
// 	q.Add("content", "Your Chative verification code is: "+vCode)
// 	req.URL.RawQuery = q.Encode()
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()
// 	d, err = ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return err
// 	}
// 	var r = struct {
// 		Code string `json:"code"`
// 	}{}
// 	err = json.Unmarshal(d, &r)
// 	if err != nil {
// 		return err
// 	}
// 	if r.Code != "SUCCESS" {
// 		return errors.New(string(d))
// 	}
// 	return nil
// }
