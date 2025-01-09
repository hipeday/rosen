import Guide from '@/components/Guide';
import { trim } from '@/utils/format';
import { PageContainer } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import styles from './index.less';
import { useState } from 'react';
import { QRCode } from 'antd';

const HomePage: React.FC = () => {
  const { name } = useModel('global');
  const [ zfaUrl, set2faUrl ] = useState<string>("otpauth://totp/Rosen%20Console%20Admin%20Panel:ggboy?algorithm=SHA1&digits=6&issuer=Rosen%20Console%20Admin%20Panel&period=30&secret=FQVHCNGTZVWNHSTL56MDJ4FL3FCXH5O3");
  return (
    <PageContainer ghost>
    
      <div className={styles.container}>
        <Guide name={trim(name)} />
      </div>

      <div>
        <QRCode value={zfaUrl} />
      </div>
      
    </PageContainer>
  );
};

export default HomePage;
