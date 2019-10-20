package net.away0x.lib_user_center.ui.activity

import android.os.Bundle
import android.util.Log
import android.view.View
import kotlinx.android.synthetic.main.activity_login.*
import kotlinx.android.synthetic.main.activity_login.mMobileEt
import kotlinx.android.synthetic.main.activity_login.mPwdEt
import net.away0x.lib_base.ext.enable
import net.away0x.lib_base.ext.onClick
import net.away0x.lib_base.ui.activity.BaseMvpActivity
import net.away0x.lib_user_center.R
import net.away0x.lib_user_center.data.protocol.UserInfo
import net.away0x.lib_user_center.injection.component.DaggerUserComponet
import net.away0x.lib_user_center.injection.module.UserModule
import net.away0x.lib_user_center.presenter.LoginPresenter
import net.away0x.lib_user_center.presenter.view.LoginView
import org.jetbrains.anko.startActivity
import org.jetbrains.anko.toast

/** 登录界面 */
class LoginActivity : BaseMvpActivity<LoginPresenter>(), LoginView,
    View.OnClickListener {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)

        initView()
    }

    override fun onLoginResult(result: UserInfo) {
        Log.d("asd", result.toString())
        toast("登录成功")
    }

    override fun injectComponent() {
        DaggerUserComponet
            .builder()
            .activityComponent(activityComponent)
            .userModule(UserModule())
            .build()
            .inject(this) // 注入

        mPresenter.mView = this
    }

    override fun onClick(view: View?) {
        when (view?.id) {
            // 跳转注册界面
            R.id.mRightTv -> {
                startActivity<RegisterActivity>()
            }

            // 登录
            R.id.mLoginBtn -> {
                mPresenter.login(mMobileEt.text.toString(), mPwdEt.text.toString(), "")
                hideInputBoard()
            }
        }
    }

    private fun initView() {
        mHeaderBar.getRightView().onClick(this)
        mLoginBtn.onClick(this)
        // 判断按钮是否可用
        mLoginBtn.enable(mMobileEt) { isRegisterBEnEnable() }
        mLoginBtn.enable(mPwdEt) { isRegisterBEnEnable() }
    }

    /**
     * 注册按钮是否可点击
     */
    private fun isRegisterBEnEnable(): Boolean {
        return mMobileEt.text.isNullOrEmpty().not() &&
                mPwdEt.text.isNullOrEmpty().not()
    }
}
