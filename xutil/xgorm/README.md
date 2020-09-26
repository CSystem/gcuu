# gorm 的辅助类

`Transact` 用来简化事务的写法。

## Example

```go
func (s *withdrawService) Create(withdraw *entity.Withdraw, uid uint, params *request.Withdraw) error {
  mutex := s.redsync.NewMutex(fmt.Sprintf("user-wallet-%d", uid))
  if err := mutex.Lock(); err != nil {
    return err
  }
  defer mutex.Unlock()

  // 这里可以调用 Transact 方法，将需要在事务中运行的过程传入
  return xgorm.Transact(s.db, func(tx *gorm.DB) error {
    var wallet entity.Wallet
    if err := tx.Where("user_id = ?", uid).First(&wallet).Error; err != nil {
      return err
    }

    if wallet.Money < params.Amount {
      return consts.ErrBalanceNotEnough
    }

    if err := tx.Model(&wallet).Where("user_id = ?", uid).UpdateColumn("money", gorm.Expr("money - ?", params.Amount)).Error; err != nil {
      return err
    }

    if err := copier.Copy(withdraw, params); err != nil {
      return err
    }

    if err := tx.Create(&withdraw).Error; err != nil {
      return err
    }
    return nil
  })
}
```
