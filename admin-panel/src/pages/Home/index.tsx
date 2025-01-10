import Guide from '@/components/Guide';
import { trim } from '@/utils/format';
import { PageContainer } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import styles from './index.less';
import { useState } from 'react';
import { QRCode } from 'antd';

const HomePage: React.FC = () => {
  const { name } = useModel('global');
  const [ zfaUrl ] = useState<string>("otpauth://totp/Rosen%20Console%20Admin%20Panel:ggboy?algorithm=SHA1&digits=6&issuer=Rosen%20Console%20Admin%20Panel&period=30&secret=FQVHCNGTZVWNHSTL56MDJ4FL3FCXH5O3");
  const [ captcha, setCaptcha ] = useState<string>('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPAAAABQCAMAAAAQlwhOAAAA81BMVEUAAABqDlTihszTd72HK3GSNnzafsTbf8WNMXfpjdORNXvJbbNwFFpeAkhzF13kiM7ihsxgBEqiRoxeAkiJLXOSNnyoTJLQdLqnS5HKbrSdQYfNcbetUZe1WZ+PM3nLb7WUOH7egsiUOH7ni9HPc7lyFlzVeb+bP4WGKnBeAkjOcrjghMq6XqSXO4HpjdO5XaN+ImixVZvTd72PM3ndgcfZfcOnS5GtUZefQ4m5XaOtUZekSI50GF5wFFq3W6FjB021WZ9wFFqdQYerT5WeQojCZqyCJmzjh83wlNrKbrTtkdd2GmBzF11mClCYPILTd72mSpCJUGRMAAAAAXRSTlMAQObYZgAABmBJREFUeJzsXAtLIzEQnlFQoYInqK2goKjoVUFRsIjig6JYqf//7xzuI5mZTLLJPmq19x3YZjfZmS9fHpNkezBXGANsfrcPs8R4PN7cXCzGC6bwIuB9JlaelGsXM7EM8MxS7++zYPz05DK+uJgN4+dnwZh8n8Y8YDvd5k53Ct9VZyn5Xjt3plONMQLS5PZ2MuOdnZ3UIrG4u4tgnOP6WmFsvqH994Xie4ZaCneGaL6awgCwm3+gjra8nBMg4O5uzvhb+CLK7hPCXkNjRePVFC7bdWPCW37z1kiklb29BoxpR7XmgfbchKr3YmvLx5iNFLGM/beuQuWILVu/sqpbatFevsgUdg31k8xcXfkZM7JSYQCYiHRHyB0wHUva6vcTGQPAW/71URhyyRbXcz0nkwlLdwXG0W3V2E+2/vaWMX585IwtW27eWJzY9EvQwCTVIxAmuXUU7qQ3sVxhFHz1HiP5Z/Xy8mIYL7slyqbQBpA3KNYKE59UPmdInx1VjCq8vMwZZ09rjy9v4MgH1cqyS9pzhsNheSWKr6wXwXc4jCDBMagyiYSxZ5xRsbRkGWNDhb0Ypk4bMBiEGEuBgYzfqPQ2dmWpnIS1mTRe4AqkThsVCguBTcJpajYUtFfNJKwFMdECV+VJ5FtlEFC3LkYusp4jSl+ZrGGFfaRqDI5NIUmSO/QeDfYlO0+YWl4JjIDxsW0IxymZhcLAY1zgrd0ozMT3LERkWKnNyFEh1kbF/ePjeMboTMOoO2QpOSst70JELBxgTbEfofDGRsnYNxhlfD+qnmN5cA/EXeuZ059JQtEK5cphbU0yDs5Itg9tFB1pMBiQ3KP847ZIfnzEMA7FuSzp9loq8BmqMbmUeE3L4HrkWMPyz4B4NBpljG9vDWOW3cPXFcaNM9WaYOKfnZ1pszYJUNXYze34+RqOLZrRDhkmd/4pFAZRQSpfJWKm11D2Yelt+e1MbZ7uijCscCGuonAeC9nn6Pi0j1WzeIZOJGMQyp6q0sEqT4wbisIlU+R/ZFnXRSfb56dlLG2dBydH0/PQE4Np6eoBVy5DS/s5SYDUGflJFvikptiMc35+7l2r5tWjbCQiaRJI6ph2ubCHcv7Dsql6dA0/LDtOQsSHkK3S5Dl4FoBH3l0Zuz4o/1g37dgCgcBXVbgoVSfcyhV+eNAZWymMu4rCR0dHbPY0npbjRpXCiNjv99UuBkqcCuFJpBpfCofahpVFGbC+hvkjnNp6xyIOA8ES2QThKmxZ5XRQGDcBm1otBFFHRpW9305pgu9oBIjT6dShJRSWhbcJHzLaFRVF0ohKoBriG3coSCctXwZ2blfmHWXpKSLQ4ApFF3MIm+NN86gYhfkI8dfDOIZvYamiXTOXDW17RypiKDnr4VJhX6eVD6R3DN+/HsYpqCAMfP42VX+Bgh+6/uudr7qS5VaSvd2cb2gN4rhmGzay9yIuuZfAe6Y2af0JeOTUXSqnCgQD6LFvRxIvrCOXl5falgd4FvxffP/4GIvcne/rsNdMEMbjMSqRhez9lyBCZvSmirJehbXTjQ4psxdrEHu9MfpOD9U1HnFSiQ3ZFZLidcOLdnxaKBXu8SUR9c0ZetCzQFx3YlN0loPINwGoCw1AT6qj3oMC4PtS7hTFcnKJy8vr6+vaslWO8bZmWxJ0hb2LoL8H5QdRmDQ/ZfKV3zKF5RVkK2Q7hykrydpYWVmppzBxM6yTT2HtimnEtjErzbkZVtKy3+iXlfjCeTlBnTxFb5AEWyabjpsbD2MGZ7PR3FAl1guVV1uA/2WkasTwBSYWn07kzFtsSTs11Kaq/peR2gTzX1m15+meWP94opmGmAVf0P23+27Zv16vZ1PdkO0YbyylTZ9M+ExhpCvcH4bi3SMGd41rd4hMn/6BXHO4fDX4tgd+LxCVtdKvxsIQ/TU4/G4HZozDw4Vj/N0O/Da0+l7ZXOAzeDf5zcF5wb3vhj171zCcD4VPk0vc3/sZB4rVef23A5ye1mCclv2k+JwLvnUUTsTJyUnXJuYM88D3tUHZ8K9O5hOvr/UZk1+dfCMCx5MqNL6xZ7Mu3/3sb8IPbFOwrl30H0/Go/7p+/7+ftpPqFOQHeW4yPmuNnp0/dP31hU+IN8De5Wrq4Sx9uP/n4KDA8s4uBtN+Sr/vcPPQaTCDD+Z73/8RxT+BQAA//9iY0hJXZ82TgAAAABJRU5ErkJggg==');
  return (
    <PageContainer ghost>
    
      <div className={styles.container}>
        <Guide name={trim(name)} />
      </div>

      <div>
        <QRCode value={zfaUrl} />
      </div>

      <div>
        <img onClick={async () => {
          const { captcha } = await (await fetch('/api/console/captcha')).json();
          setCaptcha(captcha);
        }} src={captcha} alt="验证码" />
      </div>
      
    </PageContainer>
  );
};

export default HomePage;
