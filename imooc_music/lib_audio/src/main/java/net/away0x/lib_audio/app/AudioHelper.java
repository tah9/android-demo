package net.away0x.lib_audio.app;

import android.content.Context;

import net.away0x.lib_audio.mediaplayer.core.AudioController;
import net.away0x.lib_audio.mediaplayer.core.MusicService;
import net.away0x.lib_audio.mediaplayer.db.GreenDaoHelper;
import net.away0x.lib_audio.mediaplayer.model.AudioBean;

import java.util.ArrayList;

/**
 * @function 唯一与外界通信的帮助类
 */
public final class AudioHelper {

    //SDK全局Context, 供子模块用
    private static Context mContext;

    public static void init(Context context) {
        mContext = context;
        GreenDaoHelper.initDatabase();
    }


    public static Context getContext() {
        return mContext;
    }

    //外部启动MusicService方法
    public static void startMusicService(ArrayList<AudioBean> audios) {
        MusicService.startMusicService(audios);
    }

    public static void pauseAudio() {
        AudioController.getInstance().pause();
    }

    public static void resumeAudio() {
        AudioController.getInstance().resume();
    }
}